// Пустой шаблон проекта на Golang
package main

import (
	"errors"
	"log"
	"project/internal/config"
	"project/pkg/logger"

	flag "github.com/spf13/pflag"
)

func main() {
	var cfgPath string
	flag.StringVarP(&cfgPath, "config", "c", "config", "path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	zerolog := logger.New(cfg.Log.Level, cfg.Log.File)
	zerolog.Info().Msg("Empty project template on Golang")
	zerolog.Info().Str("Version", config.ShowVersion()).Send()
	zerolog.Debug().Msgf("%#v", cfg)
	zerolog.Error().Stack().Err(errors.New("file open failed")).Msg("something happened!")
	zerolog.Info().Msg("the end")
}
