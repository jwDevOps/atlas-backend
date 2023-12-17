package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Name string `yaml:"name"`
}

func Load() *Configuration {
	f, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var c Configuration
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	return &c
}
