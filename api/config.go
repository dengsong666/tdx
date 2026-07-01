package main

import "os"

type Config struct {
	Addr string
}

func LoadConfig() Config {
	addr := os.Getenv("TDX_API_ADDR")
	if addr == "" {
		addr = ":8181"
	}
	return Config{Addr: addr}
}
