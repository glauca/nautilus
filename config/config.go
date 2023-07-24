package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("no such config file")
		} else {
			log.Println("read config error")
		}
		log.Fatal(err)
	}
}
