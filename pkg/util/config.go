package util

import "flag"

type Config struct {
	Addr int `json:"addr"`
}

func NewConfig() *Config {
	cfg := &Config{}
	flag.IntVar(&cfg.Addr, "addr", 8080, "HTTP network address")
	flag.Parse()
	return cfg
}
