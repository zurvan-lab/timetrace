package errors

import "errors"

var (
	// config.
	ErrInavlidConfigPath               = errors.New("invalid config path")
	ErrInvalidUsers                    = errors.New("invalid user(s)")
	ErrSpecificAndAllCommandSameAtTime = errors.New("can't have all cmds and specific cmd at same time")

	// server.
	ErrAuth = errors.New("authentication error")

	// CLI.
	ErrInvalidUserOrPassword = errors.New("user or user information you provided is invalid")
)
