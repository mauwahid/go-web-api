package config

import "time"

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
