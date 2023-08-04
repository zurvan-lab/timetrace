package src

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Authorization bool
	Proc          struct {
		Cores, Threads int
	}
	Listen struct {
		Ip, Port string
	}
	Log struct {
		Path string
	}
	User struct {
		Name string
		Token string
	}
}

func createConfig() *Config {
	return &Config{}
}

func ReadConfigFile(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		// TODO: log
	}
	defer file.Close()

	config := createConfig()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		// TODO: log
	}
	//TODO: validate config
	//TODO: Log
	return config
}
