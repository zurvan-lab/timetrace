package database

import "time"

func NewElement(v string, t time.Time) Element {
	return Element{value: v, time: t}
}
