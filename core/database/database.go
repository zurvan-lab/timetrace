package database

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/zurvan-lab/timetrace/config"
)

type Database struct {
	Sets   Sets
	Config *config.Config

	sync.RWMutex
}

func Init(cfg *config.Config) *Database {
	return &Database{
		Sets:   make(Sets, 1024),
		Config: cfg,
	}
}

func (db *Database) SetsMap() Sets {
	return db.Sets
}

// ! TQL Commands.
func (db *Database) Connect(args []string) string {
	db.RLock()
	defer db.RUnlock()

	if len(args) != 2 {
		return INVALID
	}

	for _, u := range db.Config.Users {
		if u.Name == args[0] && u.Password == args[1] {
			return OK
		}
	}

	return INVALID
}

func (db *Database) Ping(_ []string) string {
	return PONG
}

func (db *Database) AddSet(args []string) string {
	db.Lock()
	defer db.Unlock()

	if len(args) != 1 {
		return INVALID
	}

	db.Sets[args[0]] = make(Set) // args[0] is set name. see: TQL docs.

	return OK
}

func (db *Database) AddSubSet(args []string) string {
	db.Lock()
	defer db.Unlock()

	if len(args) != 2 {
		return INVALID
	}

	s, ok := db.Sets[args[0]] // set name args[0]
	if !ok {
		return SET_NOT_FOUND
	}

	s[args[1]] = make(SubSet, 0) // subset name args[1]

	return OK
}

func (db *Database) PushElement(args []string) string {
	db.Lock()
	defer db.Unlock()

	if len(args) != 4 {
		return INVALID
	}

	setName := args[0]
	subSetName := args[1]
	elementValue := args[2]
	timeStr := args[3]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return SUB_SET_NOT_FOUND
	}

	timestamp, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return INVALID
	}

	t := time.Unix(timestamp, 0)
	e := NewElement(elementValue, t)

	db.Sets[setName][subSetName] = append(db.Sets[setName][subSetName], e)

	return OK
}

func (db *Database) DropSet(args []string) string {
	db.Lock()
	defer db.Unlock()

	if len(args) != 1 {
		return INVALID
	}

	setName := args[0]
	_, ok := db.Sets[setName]

	if !ok {
		return SET_NOT_FOUND
	}

	delete(db.Sets, setName)

	return OK
}

func (db *Database) DropSubSet(args []string) string {
	db.Lock()
	defer db.Unlock()

	if len(args) != 2 {
		return INVALID
	}

	setName := args[0]
	subSetName := args[1]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return SUB_SET_NOT_FOUND
	}

	delete(db.Sets[setName], subSetName)

	return OK
}

func (db *Database) CleanSets(_ []string) string {
	db.Lock()
	defer db.Unlock()

	db.Sets = make(Sets)

	return OK
}

func (db *Database) CleanSet(args []string) string {
	db.Lock()
	defer db.Unlock()

	if len(args) != 1 {
		return INVALID
	}

	setName := args[0]

	_, ok := db.Sets[setName]
	if !ok {
		return SET_NOT_FOUND
	}

	db.Sets[setName] = make(Set)

	return OK
}

func (db *Database) CleanSubSet(args []string) string {
	db.Lock()
	defer db.Unlock()

	if len(args) != 2 {
		return INVALID
	}

	setName := args[0]
	subSetName := args[1]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return SUB_SET_NOT_FOUND
	}

	db.Sets[setName][subSetName] = make(SubSet, 0)

	return OK
}

func (db *Database) CountSets(_ []string) string {
	db.RLock()
	defer db.RUnlock()

	return fmt.Sprint(len(db.Sets))
}

func (db *Database) CountSubSets(args []string) string {
	db.RLock()
	defer db.RUnlock()

	if len(args) != 1 {
		return INVALID
	}

	subSet, ok := db.Sets[args[0]]
	if !ok {
		return SET_NOT_FOUND
	}

	return fmt.Sprint(len(subSet))
}

func (db *Database) CountElements(args []string) string {
	db.RLock()
	defer db.RUnlock()

	if len(args) != 2 {
		return INVALID
	}

	elms, ok := db.Sets[args[0]][args[1]]
	if !ok {
		return SUB_SET_NOT_FOUND
	}

	return fmt.Sprint(len(elms))
}

func (db *Database) GetElements(args []string) string {
	db.RLock()
	defer db.RUnlock()

	if len(args) < 2 {
		return INVALID
	}

	subSet, ok := db.Sets[args[0]][args[1]]
	if !ok {
		return SUB_SET_NOT_FOUND
	}

	if len(args) == 3 {
		n, err := strconv.Atoi(args[2])
		if err != nil || len(subSet) < n {
			return INVALID
		}

		lastN := subSet[len(subSet)-n:]

		return lastN.String()
	}

	return subSet.String()
}
