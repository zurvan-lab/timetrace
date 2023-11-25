package database

import (
	"strconv"
	"time"

	"github.com/zurvan-lab/TimeTrace/config"
)

type Database struct {
	Sets   Sets
	Config *config.Config
}

func Init(path string) IDataBase {
	return &Database{
		Sets:   make(Sets),
		Config: config.LoadFromFile(path),
	}
}

func (db *Database) SetsMap() Sets {
	return db.Sets
}

// ! Commands.
func (db *Database) AddSet(args []string) string {
	db.Sets[args[0]] = make(Set) // args[0] is set name. see: TQL docs.

	return "DONE"
}

func (db *Database) AddSubSet(args []string) string {
	s, ok := db.Sets[args[0]] // set name args[0]
	if !ok {
		return "SNF"
	}

	s[args[1]] = make(SubSet, 0) // subset name args[1]

	return "DONE"
}

func (db *Database) PushElement(args []string) string {
	setName := args[0]
	subSetName := args[1]
	elementValue := []byte(args[2])
	timeStr := args[3]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return "SSNF"
	}

	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return "INVALID"
	}

	t := time.Unix(timestamp, 0)
	e := NewElement(elementValue, t)

	db.Sets[setName][subSetName] = append(db.Sets[setName][subSetName], e)

	return "DONE"
}

func (db *Database) DropSet(args []string) string {
	setName := args[0]
	_, ok := db.Sets[setName]

	if !ok {
		return "SNF"
	}

	delete(db.Sets, setName)

	return "DONE"
}

func (db *Database) DropSubSet(args []string) string {
	setName := args[0]
	subSetName := args[1]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return "SSNF"
	}

	delete(db.Sets[setName], subSetName)

	return "DONE"
}

func (db *Database) CleanSets(_ []string) string {
	db.Sets = make(Sets)

	return "DONE"
}

func (db *Database) CleanSet(args []string) string {
	setName := args[0]

	_, ok := db.Sets[setName]
	if !ok {
		return "SNF"
	}

	db.Sets[setName] = make(Set)

	return "DONE"
}

func (db *Database) CleanSubSet(args []string) string {
	setName := args[0]
	subSetName := args[1]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return "SSNF"
	}

	db.Sets[setName][subSetName] = make(SubSet, 0)

	return "DONE"
}
