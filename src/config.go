package src

import (
	"os"

	"github.com/zurvan-lab/TimeTraceDB/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Authorization bool
	Proc          struct {
		Cores, Threads int
	}
	Listen struct {
		IP, Port string
	}
	Log struct {
		Path string
	}
	User struct {
		Name  string
		Token string
	}
}

func createConfig() *Config {
	return &Config{}
}

func ReadConfigFile(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Error("Can not open the config file", "error: ", err)
	}
	defer file.Close()

	config := createConfig()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Error("Can not decode the Config Yaml file", "error: ", err)
	}
	//TODO: validate config
	//TODO: Log
	return config
}
