package server

import "fmt"

type AuthenticateError struct {
	reason string
}

func (e AuthenticateError) Error() string {
	return fmt.Sprintf("auth error: %s", e.reason)
}
