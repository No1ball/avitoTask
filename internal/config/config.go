package config

import "github.com/spf13/viper"

func InitConfig() error {
	viper.AddConfigPath("../avitoEx/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
