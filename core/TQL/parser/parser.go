package core

import (
	"strings"

	"github.com/zurvan-lab/timetrace/core/database"
)

// parsing TQL queries. see: docs/TQL.
func ParseQuery(query string) database.Query {
	q := strings.Split(query, " ")

	return database.Query{
		Command: q[0],
		Args:    q[1:],
	}
}
