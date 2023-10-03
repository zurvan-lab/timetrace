package database

import (
	"github.com/zurvan-lab/TimeTraceDB/config"
)

type Database struct {
	Sets   Sets
	Config *config.Config
}

func Init(path string) *Database {
	return &Database{
		Sets:   make(Sets),
		Config: config.LoadFromFile(path),
	}
}

func (db *Database) AddSet(name string) string {
	db.Sets[name] = make(Set)
	return "DONE"
}

func (db *Database) AddSubSet(set, name string) string {
	s, ok := db.Sets[set]
	if !ok {
		return "SETNF"
	}
	s[name] = make(SubSet, 0)
	return "DONE"
}

func (db *Database) PushElement(set, subset string, e Element) string {
	_, ok := db.Sets[set][subset]
	if !ok {
		return "SUBSETNF"
	}
	db.Sets[set][subset] = append(db.Sets[set][subset], e)
	return "DONE"
}

func (db *Database) DropSet(name string) string {
	_, ok := db.Sets[name]
	if !ok {
		return "SETNF"
	}
	delete(db.Sets, name)
	return "DONE"
}

func (db *Database) DropSubSet(set, subset string) string {
	_, ok := db.Sets[set][subset]
	if !ok {
		return "SUBETNF"
	}
	delete(db.Sets[set], subset)
	return "DONE"
}

func (db *Database) CleanSets() string {
	db.Sets = make(Sets)
	return "DONE"
}

func (db *Database) CleanSet(name string) string {
	_, ok := db.Sets[name]
	if !ok {
		return "SETNF"
	}
	db.Sets[name] = make(Set)
	return "DONE"
}

func (db *Database) CleanSubSet(set, subset string) string {
	_, ok := db.Sets[set][subset]
	if !ok {
		return "SUBETNF"
	}
	db.Sets[set][subset] = make(SubSet, 0)
	return "DONE"
}
