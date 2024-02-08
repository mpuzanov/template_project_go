package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	// Version для установки версии релиза
	Version = "development"
)

// ShowVersion Вывод версии релиза программы
func ShowVersion() string {
	return Version
}

// Config ...
type Config struct {
	Env  string     `yaml:"env" env-default:"local"`
	Mail MailConfig `yaml:"mail"`
	DB   DBConf     `yaml:"db"`
}

// DBConf .
type DBConf struct {
	Name     string `yaml:"name" env:"DB_NAME"`
	Host     string `yaml:"host" env:"DB_HOST" env-required:"true"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"1433" env-description:"sql server port"`
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Database string `yaml:"database" env:"DB_DATABASE"`
}

// MailConfig ...
type MailConfig struct {
	Server      string `yaml:"server" env:"mail_server"`
	Port        int    `yaml:"port" env:"mail_port"`
	UseTLS      bool   `yaml:"use_tls" env:"mail_use_tls" env-default:"false"`
	UseSSL      bool   `yaml:"use_ssl" env:"mail_use_ssl" env-default:"true"`
	UserName    string `yaml:"username" env:"mail_username"`
	Password    string `yaml:"password" env:"mail_password"`
	Timeout     int    `yaml:"timeout" env:"mail_timeout" env-default:"10"`
	IsSendEmail bool   `yaml:"is_send_email"`
}

// MustConfig ...
func MustConfig(fileConf string) *Config {

	instance := &Config{}

	if _, err := os.Stat(fileConf); errors.Is(err, os.ErrNotExist) {
		workingDir, _ := os.Executable()
		fileConf = filepath.Join(filepath.Dir(workingDir), filepath.Base(fileConf))
	}

	if err := cleanenv.ReadConfig(fileConf, instance); err != nil {
		log.Fatal(err)
	}

	return instance
}
