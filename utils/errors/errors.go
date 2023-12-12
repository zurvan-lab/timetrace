package errors

import "errors"

var (
	ErrInavlidConfigPath               = errors.New("invalid config path")
	ErrInvalidUsers                    = errors.New("invalid user(s)")
	ErrSpecificAndAllCommandSameAtTime = errors.New("can't have all cmds and specific cmd at same time")
)
