package errors

import "errors"

var (
	// Config errors.
	ErrInavlidConfigPath               = errors.New("invalid config path")
	ErrInvalidUsers                    = errors.New("invalid user(s)")
	ErrSpecificAndAllCommandSameAtTime = errors.New("can't have all cmds and specific cmd at same time")
	ErrEmptyLogTarget                  = errors.New("log target can't be empty when the logger is enabled")

	// Server errors.
	ErrAuth = errors.New("authentication error")

	// CLI errors.
	ErrInvalidUserOrPassword = errors.New("user or user information you provided is invalid")
	ErrInvalidCommand        = errors.New("invalid command")
)
