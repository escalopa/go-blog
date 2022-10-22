package main

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
}

var AppConfig Config

func LoadConfiguration() {
	log.Println("Loading Configuration Data")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
