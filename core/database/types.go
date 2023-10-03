package database

import "time"

type Element struct {
	value string
	time  time.Time
}

type (
	Sets   map[string]Set
	Set    map[string]SubSet
	SubSet []Element
)
