package config

import "flag"

type Config struct {
	ServerPort string
}

func (c *Config) EnableFlags(isEnable bool) {
	if isEnable {
		flag.StringVar(&c.ServerPort, "server-port", "8080", "Put your custom server port here!")
		flag.Parse()
	} else {
		c = &Config{}
	}
}
