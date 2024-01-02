package database

import (
	"fmt"
	"strings"
	"time"
)

const (
	INVALID           = "INVALID"
	OK                = "OK"
	PONG              = "PONG"
	SET_NOT_FOUND     = "SNF"
	SUB_SET_NOT_FOUND = "SSNF"
	ELEMENT_NOT_FOUND = "ENF"
)

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

func (ss SubSet) String() string {
	var builder strings.Builder
	for _, s := range ss {
		builder.WriteString(fmt.Sprintf(" %v-%d ", s.value, s.time.Unix()))
	}

	return builder.String()
}
