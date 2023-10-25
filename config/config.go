package config

import (
	"bytes"
	_ "embed"
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configBytes []byte

type Config struct {
	Name   string `yaml:"name"`
	Server Server `yaml:"server"`
	Log    Log    `yaml:"log"`
	Users  []User `yaml:"users"`
}

type Server struct {
	IP   string `yaml:"listen"`
	Port string `yaml:"port"`
}

type Log struct {
	Path string `yaml:"path"`
}

type User struct {
	Name     string   `yaml:"name"`
	Password string   `yaml:"password"`
	Cmds     []string `yaml:"cmds"`
}

func (conf *Config) BasicCheck() error {
	if len(conf.Users) <= 0 {
		return errors.New("invalid user length")
	}
	return nil
}

func DefaultConfig() *Config {
	config := &Config{
		Server: Server{
			IP: "localhost",
			Port: "7070",
		},
		Log: Log{
			Path: "ttrace.log",
		},
		Name: "time_trace",
	}
	rootUser := User{
		Name:     "root",
		Password: "super_secret_password",
		Cmds:     []string{"*"},
	}
	config.Users = append(config.Users, rootUser)

	return config
}

func LoadFromFile(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	config := &Config{}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	err = config.BasicCheck()
	if err != nil {
		panic(err)
	}

	return config
}

func (conf *Config) ToYAML() []byte {
	buf := new(bytes.Buffer)
	encoder := yaml.NewEncoder(buf)
	err := encoder.Encode(conf)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}
