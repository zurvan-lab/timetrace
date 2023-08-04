package src

import (
	"net"
	"os"
)

type Database struct {
	Sets       Sets
	Config     Config
	socket     net.Listener
	Users      Users
}

func CreateDataBase() *Database {
	return &Database{
		Sets: *NewSets(),
		Config: *ReadConfigFile(""),
	}
}

func (db *Database)InitSocket()  {
	var err error
	db.socket ,err = net.Listen("tcp",db.Config.Listen.Ip+":"+db.Config.Listen.Port)
	if err != nil {
		os.Exit(1)
	}
}

func (db *Database) InitUsers()  {
	users:= CreateUsers()
	db.Users = *users
	cmds:= []string{"all"}
	users.NewUser(db.Config.User.Name,db.Config.User.Token,cmds)
}
