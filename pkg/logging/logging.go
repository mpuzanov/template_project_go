package logging

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type (
	Logger = slog.Logger
	Attr   = slog.Attr
)

var (
	StringAttr   = slog.String
	BoolAttr     = slog.Bool
	Float64Attr  = slog.Float64
	AnyAttr      = slog.Any
	DurationAttr = slog.Duration
	IntAttr      = slog.Int
	Int64Attr    = slog.Int64

	GroupValue = slog.GroupValue
	Group      = slog.Group
)

// NewLogger ...
func NewLogger(env string) *Logger {
	var log *Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envProd:
		// добавим запись логов в файл
		ex, _ := os.Executable()
		fileName := filepath.Base(ex)
		fileLog := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".log"
		workDir := filepath.Dir(ex) // путь к программе
		fileLog = filepath.Join(workDir, fileLog)

		file := &lumberjack.Logger{
			Filename:   fileLog, // Имя файла
			MaxSize:    10,      // Размер в МБ до ротации файла
			MaxBackups: 5,       // Максимальное количество файлов, сохраненных до перезаписи
			MaxAge:     30,      // Максимальное количество дней для хранения файлов
			Compress:   true,    // Следует ли сжимать файлы логов с помощью gzip
		}

		multi := io.MultiWriter(file, os.Stdout) //, os.Stderr

		log = slog.New(
			slog.NewJSONHandler(multi, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

//=======================================================

// ErrAttr helper func
// logger.Error("user msg error", ErrAttr(err))
func ErrAttr(err error) Attr {
	return slog.String("error", err.Error())
}

//=======================================================

type ctxLogger struct{}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*slog.Logger); ok {
		return l
	}
	return slog.Default()
}
