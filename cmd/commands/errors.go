package commands

import "fmt"

type InvalidAuthInfoError struct {
	command string
}

func (e InvalidAuthInfoError) Error() string {
	return fmt.Sprintf("command %s is not a valid command to authenticate", e.command)
}

type InvalidConfigPathError struct {
	path string
}

func (e InvalidConfigPathError) Error() string {
	return fmt.Sprintf("path %s doesn't contain a valid config file", e.path)
}
