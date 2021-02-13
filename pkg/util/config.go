package util

import "flag"

type Config struct {
	HttpPort   int    `json:"port"`
	DbEnable   bool   `json:"enable"`
	DbUser     string `json:"user"`
	DbPassword string `json:"password"`
	DbHost     string `json:"host"`
}

func NewConfig() *Config {
	cfg := &Config{}
	flag.IntVar(&cfg.HttpPort, "port", 8080, "HTTP network port")
	flag.BoolVar(&cfg.DbEnable, "enable", false, "MySQL enable")
	flag.StringVar(&cfg.DbUser, "user", "root", "MySQL user")
	flag.StringVar(&cfg.DbPassword, "password", "123", "MySQL password")
	flag.StringVar(&cfg.DbHost, "host", "localhost:3306", "MySQL host")
	flag.Parse()
	return cfg
}
