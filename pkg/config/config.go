package config

import (
	"github.com/spf13/viper"
)

var (
	Cfg *Config
)

type MysqlConfig struct {
	Host     string
	Port     int32
	User     string
	Password string
	Database string
}

type Config struct {
	Mysql MysqlConfig
}

func init() {
	cfg := Config{}

	viper.SetConfigFile("config/config.toml")
	viper.ReadInConfig()
	viper.Unmarshal(&cfg)
	Cfg = &cfg
}
