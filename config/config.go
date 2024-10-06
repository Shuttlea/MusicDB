package config

import (
	"log/slog"

	"github.com/gofor-little/env"
)

func InitConfig() *Config {
	cfg := Config{}
	err := env.Load("config/config.env")
	if err != nil {
		slog.Debug(err.Error())
		panic(err)
	}
	slog.Debug("Config file loaded")
	cfg.HostAddr, err = env.MustGet("HOST_ADDR")
	if err != nil {
		slog.Debug(err.Error())
	}
	cfg.HostDB, err = env.MustGet("DATABASE_HOST")
	if err != nil {
		slog.Debug(err.Error())
	}
	cfg.UserDB, err = env.MustGet("DATABASE_USER")
	if err != nil {
		slog.Debug(err.Error())
	}
	// cfg.PasswordDB,err = env.MustGet("DATABASE_PASSWORD")
	// if err!= nil{
	// 	slog.Debug(err.Error())
	// }
	cfg.DataBase, err = env.MustGet("DATABASE_DB")
	if err != nil {
		slog.Debug(err.Error())
	}
	cfg.PortDB, err = env.MustGet("DATABASE_PORT")
	if err != nil {
		slog.Debug(err.Error())
	}
	cfg.SslmodeDB, err = env.MustGet("DATABASE_SSLMODE")
	if err != nil {
		slog.Debug(err.Error())
	}
	return &cfg
}

type Config struct {
	HostAddr   string `env:"HOST_ADDR"`
	HostDB     string `env:"DATABASE_HOST"`
	UserDB     string `env:"DATABASE_USER"`
	PasswordDB string `env:"DATABASE_PASSWORD"`
	DataBase   string `env:"DATABASE_DB"`
	PortDB     string `env:"DATABASE_PORT"`
	SslmodeDB  string `env:"DATABASE_SSLMODE"`
}
