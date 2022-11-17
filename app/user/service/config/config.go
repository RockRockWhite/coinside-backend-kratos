package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("./config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

type Config struct {
}

func (Config) GetString(key string) string {
	return viper.GetString(key)
}

func (Config) GetInt(key string) int {
	return viper.GetInt(key)
}

func NewConfig() *Config {
	return &Config{}
}
