package config

import "time"

type Config struct {
	AppName string
	Http    Http
}

type Http struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
