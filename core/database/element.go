package database

import "time"

type Element struct {
	value string
	time  time.Time
}

func NewElement(v string) Element {
	return Element{value: v, time: time.Now()}
}
