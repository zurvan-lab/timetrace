package database

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/zurvan-lab/TimeTrace/config"
)

type Database struct {
	Sets   Sets
	Config *config.Config

	lk sync.RWMutex
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
func (db *Database) AddSet(args []string) string {
	db.lk.Lock()
	defer db.lk.Unlock()

	if len(args) != 1 {
		return "INVALID"
	}

	db.Sets[args[0]] = make(Set) // args[0] is set name. see: TQL docs.

	return "DONE"
}

func (db *Database) AddSubSet(args []string) string {
	db.lk.Lock()
	defer db.lk.Unlock()

	if len(args) != 2 {
		return "INVALID"
	}

	s, ok := db.Sets[args[0]] // set name args[0]
	if !ok {
		return "SNF"
	}

	s[args[1]] = make(SubSet, 0) // subset name args[1]

	return "DONE"
}

func (db *Database) PushElement(args []string) string {
	db.lk.Lock()
	defer db.lk.Unlock()

	if len(args) != 4 {
		return "INVALID"
	}

	setName := args[0]
	subSetName := args[1]
	elementValue := args[2]
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
	db.lk.Lock()
	defer db.lk.Unlock()

	if len(args) != 1 {
		return "INVALID"
	}

	setName := args[0]
	_, ok := db.Sets[setName]

	if !ok {
		return "SNF"
	}

	delete(db.Sets, setName)

	return "DONE"
}

func (db *Database) DropSubSet(args []string) string {
	db.lk.Lock()
	defer db.lk.Unlock()

	if len(args) != 2 {
		return "INVALID"
	}

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
	db.lk.Lock()
	defer db.lk.Unlock()

	db.Sets = make(Sets)

	return "DONE"
}

func (db *Database) CleanSet(args []string) string {
	db.lk.Lock()
	defer db.lk.Unlock()

	if len(args) != 1 {
		return "INVALID"
	}

	setName := args[0]

	_, ok := db.Sets[setName]
	if !ok {
		return "SNF"
	}

	db.Sets[setName] = make(Set)

	return "DONE"
}

func (db *Database) CleanSubSet(args []string) string {
	db.lk.Lock()
	defer db.lk.Unlock()

	if len(args) != 2 {
		return "INVALID"
	}

	setName := args[0]
	subSetName := args[1]

	_, ok := db.Sets[setName][subSetName]
	if !ok {
		return "SSNF"
	}

	db.Sets[setName][subSetName] = make(SubSet, 0)

	return "DONE"
}

func (db *Database) CountSets(_ []string) string {
	db.lk.RLock()
	defer db.lk.RUnlock()

	i := 0
	for range db.Sets {
		i++
	}

	return fmt.Sprint(i)
}

func (db *Database) CountSubSets(args []string) string {
	db.lk.RLock()
	defer db.lk.RUnlock()

	if len(args) != 1 {
		return "INVALID"
	}

	set, ok := db.Sets[args[0]]
	if !ok {
		return "SNF"
	}

	i := 0
	for range set {
		i++
	}

	return fmt.Sprint(i)
}

func (db *Database) CountElements(args []string) string {
	db.lk.RLock()
	defer db.lk.RUnlock()

	if len(args) != 2 {
		return "INVALID"
	}

	subSet, ok := db.Sets[args[0]][args[1]]
	if !ok {
		return "SSNF"
	}

	i := 0
	for range subSet {
		i++
	}

	return fmt.Sprint(i)
}

func (db *Database) GetElements(args []string) string {
	db.lk.RLock()
	defer db.lk.RUnlock()

	if len(args) < 2 {
		return "INVALID"
	}

	subSet, ok := db.Sets[args[0]][args[1]]
	if !ok {
		return "SSNF"
	}

	if len(args) == 3 {
		n, err := strconv.Atoi(args[2])
		if err != nil || len(subSet) < n {
			return "INVALID"
		}

		lastN := subSet[len(subSet)-n:]

		return lastN.String()
	}

	return subSet.String()
}
