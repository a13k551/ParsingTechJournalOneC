package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Expression string `yaml:"expression"`
	Mask       string `yaml:"mask"`
}

var instance *Config
var once sync.Once

func Get() *Config {
	
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("../../config.yml", instance); err != nil {
			panic(err)
		}
	})

	return instance
}
