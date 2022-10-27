package main

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"db_connection_string"`
}

var AppConfig Config

func LoadConfiguration() {
	log.Println("Loading Configuration Data")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	_ = viper.BindEnv("db_connection_string", "db_connection_string")
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
