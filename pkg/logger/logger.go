package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zerolog.Logger
}

func New(logLevel string, fileLog string) *Logger {

	level, err := zerolog.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		fmt.Println(err)
	}
	zerolog.SetGlobalLevel(level)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	fileWriter := &lumberjack.Logger{
		Filename:   fileLog, // Имя файла
		MaxSize:    10,      // Размер в МБ до ротации файла
		MaxBackups: 5,       // Максимальное количество файлов, сохраненных до перезаписи
		MaxAge:     30,      // Максимальное количество дней для хранения файлов
		Compress:   true,    // Следует ли сжимать файлы логов с помощью gzip
	}

	// в файл пишем в JSON формате в консоль пишем в обычном
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05",
		FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("%s", i))
		}}

	multi := zerolog.MultiLevelWriter(fileWriter, consoleWriter)
	logger := zerolog.New(multi).With().Caller().Timestamp().Logger()

	return &Logger{&logger}
}
