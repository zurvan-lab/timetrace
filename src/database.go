package src

import (
	"net"
	"os"
)

type Database struct {
	Sets   Sets
	Config Config
	socket net.Listener
	Users  Users
}

func CreateDataBase(path string) *Database {
	return &Database{
		Sets:   *NewSets(),
		Config: *ReadConfigFile(path),
	}
}

func (db *Database) InitSocket() {
	var err error
	db.socket, err = net.Listen("tcp", db.Config.Listen.IP+":"+db.Config.Listen.Port)
	if err != nil {
		os.Exit(1)
	}
}

func (db *Database) InitUsers() {
	users := CreateUsers()
	db.Users = *users
	cmds := []string{"all"}
	users.NewUser(db.Config.FirstUser.Name, db.Config.FirstUser.Token, cmds)
}
