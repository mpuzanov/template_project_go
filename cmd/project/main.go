// Пустой шаблон проекта на Golang
package main

import (
	"errors"
	"project/internal/config"
	"project/pkg/logging"

	flag "github.com/spf13/pflag"
)

func main() {
	var cfgPath string
	flag.StringVarP(&cfgPath, "config", "c", "config.yml", "path to the configuration file")
	flag.Parse()

	cfg := config.MustConfig(cfgPath)

	logger := logging.NewLogger(cfg.Env)
	logger.Info("Empty project template on Golang")
	logger.Info("Info msg", logging.StringAttr("Version", config.ShowVersion()))
	logger.Debug("Debug msg", "Env", cfg.Env, logging.AnyAttr("cfg", cfg))
	logger.Error("Error msg", logging.ErrAttr(errors.New("test error")))
	logger.Info("the end")
}
