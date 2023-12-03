package database

import "time"

type Element struct {
	value string    // currently ttrace only supports string datatype for value.
	time  time.Time // will return and input from user as unix timestamp.
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
