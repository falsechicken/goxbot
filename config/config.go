package config

import (
	"github.com/BurntSushi/toml"
	"github.com/falsechicken/glogger"
	"os"
)

type Config struct {
	Server        string
	Username      string
	Password      string
	Status        string
	StatusMessage string
	StartTLS      bool
	Debug         bool
	Session       bool
	Console       bool
}

func Load(path string) Config {
	_, err := os.Stat(path)
	if err != nil {
		glogger.LogMessage(glogger.Warning, "Config file is missing: "+path)
	}

	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		glogger.LogMessage(glogger.Error, err.Error())
		os.Exit(2)
	}
	return config
}

func generateDefaultConfig() {
}
