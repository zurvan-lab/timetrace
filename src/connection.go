package src

import "net"

type Connection struct {
	Connection net.Listener
	User       User
	Authorized bool
}
