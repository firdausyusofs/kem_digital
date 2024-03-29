package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Logger LoggerConfig
	MySQL  MySQLConfig
}

type ServerConfig struct {
	Port      string
	JWTSecret string
}

type LoggerConfig struct {
	Encoding string
	Level    string
}

type MySQLConfig struct {
	Host     string
	DBName   string
	User     string
	Password string
	Port     string
	Driver   string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
