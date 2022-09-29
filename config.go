package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string `mapstructure:"SERVER_ADDRESS"`
	DBSource   string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (config Config, error error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return
	}
	if err := viper.Unmarshal(&config); err != nil {
		return
	}
	return
}
