package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

const (
	configTOML = "config.toml"
)

type config struct {
	Username string
	Password string
}

func getConfig() *config {
	cfg := new(config)

	_, err := toml.DecodeFile(configTOML, &cfg)
	check(err)

	if cfg.Username == "" {
		usage := read("./Helper/README.md")
		log.Panicln(string(usage))
	}

	return cfg
}
