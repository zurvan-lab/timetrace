package config

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	tte "github.com/zurvan-lab/TimeTrace/utils/errors"
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
	defaultFunction, err := defaultConf.ToYAML()
	assert.NoError(t, err)

	defaultFunctionStr := string(defaultFunction)

	defaultYaml = strings.ReplaceAll(defaultYaml, "##", "")
	defaultYaml = strings.ReplaceAll(defaultYaml, "# all commands.", "")
	defaultYaml = strings.ReplaceAll(defaultYaml, "\r\n", "\n") // For Windows
	defaultYaml = strings.ReplaceAll(defaultYaml, "\n\n", "\n")
	defaultFunctionStr = strings.ReplaceAll(defaultFunctionStr, "\n\n", "\n")

	assert.Equal(t, defaultFunctionStr, defaultYaml)
}

func TestLoadFromFile(t *testing.T) {
	config, err := LoadFromFile("./config.yaml")
	assert.NoError(t, err)

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
	assert.Error(t, tte.ErrInvalidUsers, err)

	c.Users = []User{DefaultConfig().Users[0]}
	c.Users[0].Cmds = []string{"*", "GET"}
	err = c.BasicCheck()
	assert.Error(t, tte.ErrSpecificAndAllCommandSameAtTime, err)
}
