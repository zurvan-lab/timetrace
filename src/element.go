package src

import (
	"time"
)

type Value struct {
	Data interface{}
	Time time.Time
}

type Elements map[string]Value

func NewElements() *Elements {
	return &Elements{}
}

func (e *Elements) AddElement(key string, v Value) {
	(*e)[key] = v
}
