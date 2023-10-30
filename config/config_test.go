package config

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	lines := strings.Split(string(configBytes), "\n")
	defaultYaml := ""
	for _, line := range lines {
		if !(strings.HasPrefix(line, "# ") ||
			strings.HasPrefix(line, "###") ||
			strings.HasPrefix(line, "  # ") ||
			strings.HasPrefix(line, "    # ")) {
			defaultYaml += line
			defaultYaml += "\n"
		}
	}

	defaultConf := DefaultConfig()
	defaultFunction := string(defaultConf.ToYAML())

	defaultYaml = strings.ReplaceAll(defaultYaml, "##", "")
	defaultYaml = strings.ReplaceAll(defaultYaml, "# all commands.", "")
	defaultYaml = strings.ReplaceAll(defaultYaml, "\r\n", "\n") // For Windows
	defaultYaml = strings.ReplaceAll(defaultYaml, "\n\n", "\n")
	defaultFunction = strings.ReplaceAll(defaultFunction, "\n\n", "\n")

	// fmt.Println(defaultFunction)
	// fmt.Println(defaultYaml)
	assert.Equal(t, defaultFunction, defaultYaml)
}

func TestLoadFromFile(t *testing.T) {
	config := LoadFromFile("./config.yaml")
	dConfig := DefaultConfig()

	assert.Equal(t, "time_trace", config.Name)
	assert.Equal(t, []string{"*"}, config.Users[0].Cmds)
	assert.Equal(t, "7070", config.Server.Port)
	assert.Equal(t, dConfig, config)
}

func TestBasicCheck(t *testing.T) {
	c := DefaultConfig()
	err := c.BasicCheck()
	assert.NoError(t, err)

	c.Users = []User{}
	err = c.BasicCheck()
	assert.Error(t, errors.New("invalid user length"), err)

	c.Users = []User{DefaultConfig().Users[0]}
	c.Users[0].Cmds = []string{"*", "GET"}
	err = c.BasicCheck()
	assert.Error(t, errors.New("can't have all cmds and specific cmd at same time"), err)
}
