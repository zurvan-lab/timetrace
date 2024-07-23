package config

import "fmt"

type BasicCheckError struct {
	reason string
}

func (e BasicCheckError) Error() string {
	return fmt.Sprintf("config basic check failed: %s", e.reason)
}
