package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	// Version для установки версии релиза
	Version = "development"
)

// ShowVersion Вывод версии релиза программы
func ShowVersion() string {
	return Version
}

type Config struct {
	Log LogConf `mapstructure:"log"`
	DB  DBConf  `mapstructure:"db"`
}

// LogConf структура для настройки логирования
type LogConf struct {
	Level  string `mapstructure:"level"`
	File   string `mapstructure:"file"`
	IsJSON bool   `mapstructure:"is_json"`
}

// DBConf .
type DBConf struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName(path)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("log.level", "INFO") // если параметр не задан

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
