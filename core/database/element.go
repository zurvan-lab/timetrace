package database

import "time"

func NewElement(v []byte) Element {
	return Element{value: v, time: time.Now()}
}
