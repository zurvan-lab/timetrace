package src

import (
	"os"

	"github.com/zurvan-lab/TimeTraceDB/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name          string    `yaml:"name"`
	Authorization bool      `yaml:"authorization"`
	Proc          Proc      `yaml:"proc"`
	Listen        Listen    `yaml:"server"`
	Log           Log       `yaml:"log"`
	FirstUser     FirstUser `yaml:"user"`
}

type Proc struct {
	Cores   int `yaml:"cores"`
	Threads int `yaml:"threads"`
}

type Listen struct {
	IP   string `yaml:"listen"`
	Port string `yaml:"port"`
}

type Log struct {
	Path string `yaml:"path"`
}

type FirstUser struct {
	Name  string   `yaml:"name"`
	Token string   `yaml:"token"`
	Cmd   []string `yaml:"cmd"`
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
