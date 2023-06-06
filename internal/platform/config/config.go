package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	AppName    string      `json:"appName"`
	Version    string      `json:"version"`
	HttpServer Http        `json:"httpServer"`
	Mongo      MongoConfig `json:"mongo"`
}

type MongoConfig struct {
	Address         []string      `json:"address" `
	ReplicaSet      string        `json:"replicaSet"`
	DBName          string        `json:"dbName"`
	MaxPoolSize     uint64        `json:"maxPoolSize"`
	MinPoolSize     uint64        `json:"minPoolSize"`
	MaxConnIdleTime time.Duration `json:"maxConnIdleTime"`
	Username        string        `json:"username"`
	Pass            string        `json:"pass"`
	ReadPreference  string        `json:"readPreference"`
	Timeout         time.Duration `json:"timeout"`
}

type Http struct {
	Port         int           `json:"port"`
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
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
