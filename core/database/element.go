package database

import "time"

func NewElement(v string) Element {
	return Element{value: v, time: time.Now()}
}
