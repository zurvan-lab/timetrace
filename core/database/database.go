package database

import (
	"net"
	"os"

	"github.com/zurvan-lab/TimeTraceDB/config"
)

type Database struct {
	Sets   Sets
	Config config.Config
	socket net.Listener
}

func CreateDataBase(path string) *Database {
	return &Database{
		Sets:   *NewSets(),
		Config: *config.LoadFromFile(path),
	}
}

func (db *Database) InitSocket() {
	var err error
	db.socket, err = net.Listen("tcp", db.Config.Listen.IP+":"+db.Config.Listen.Port)
	if err != nil {
		os.Exit(1)
	}
}
