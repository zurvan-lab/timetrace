package core

import (
	"strings"

	"github.com/zurvan-lab/TimeTrace/core/database"
)

// parsing TQL queries. see: docs/TQL
func ParseQuery(query string) database.Query {
	command := ""
	args := []string{}

	for _, word := range strings.Split(query, " ") {
		if word == "" {
			continue
		}

		if command != "" {
			args = append(args, word)
		} else {
			command = word
		}
	}

	return database.Query{Command: command, Args: args}
}

func Execute(query database.Query, db database.Database) string {
	return ""
}
