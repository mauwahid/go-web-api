package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	AppName  string
	Http     Http
	Database Database
}

type Http struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Host     string
	Port     string
	Name     string
	Schema   string
	Username string
	Password string
	MaxCon   string
	IdleCon  int
}

func GetConfig() *Config {
	path := os.Getenv("CONFIG_PATH")
	if len(path) == 0 {
		path = "./config.json"
	}

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	appConfig := Config{}
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}

	confData, _ := json.MarshalIndent(appConfig, "", " ")
	fmt.Println(string(confData))
	return &appConfig
}
