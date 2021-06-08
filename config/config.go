package config

import (
	"os"
	"sync"
)

var once sync.Once
var instance *Config

type Config struct {
	Directory string
}

func DefaultConfig() *Config {
	c := Instance()
	c.Directory = os.Getenv("HOME") + "/.tracker"
	return c
}

func Instance() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}
