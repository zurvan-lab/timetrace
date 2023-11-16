package database

import "time"

type Element struct {
	value []byte
	time  time.Time
}

type Query struct {
	Command string
	Args    []string
}

type (
	Sets   map[string]Set
	Set    map[string]SubSet
	SubSet []Element
)
