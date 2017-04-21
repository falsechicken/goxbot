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

//Load loads the toml config file at the path provided. If the file does not exist a default one will be created at that path.
func Load(path string) Config {
	_, err := os.Stat(path)
	if err != nil {
		glogger.LogMessage(glogger.Warning, "Config file is missing: "+path)
		generateDefaultConfig(path)
	}

	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		glogger.LogMessage(glogger.Error, err.Error())
		os.Exit(2)
	}
	return config
}

func generateDefaultConfig(path string) {
	defConf := new(Config)

	config, err := os.Create(path)
	if err != nil {
		glogger.LogMessage(glogger.Error, "Cannot create config file!: "+path)
		panic(err)
	}
	defer config.Close()

	encoder := toml.NewEncoder(config)

	defConf.Server = "example.net"
	defConf.Username = "user@example.net"
	defConf.Password = "superSecretSauce"
	defConf.Status = "a"
	defConf.StatusMessage = "GoXbot Bot"
	defConf.StartTLS = true
	defConf.Debug = false
	defConf.Session = false
	defConf.Console = false

	encoder.Encode(defConf)
}
