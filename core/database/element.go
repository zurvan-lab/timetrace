package database

import "time"

func NewElement(v []byte, t time.Time) Element {
	return Element{value: v, time: t}
}
