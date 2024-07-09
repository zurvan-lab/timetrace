package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQuery(t *testing.T) {
	t.Run("good query", func(t *testing.T) {
		query := "PUSH testSet testSubSet hello 1700842078"
		paredQuery := ParseQuery(query)

		assert.Equal(t, paredQuery.Command, "PUSH")
		assert.Equal(t, paredQuery.Args[0], "testSet")
		assert.Equal(t, paredQuery.Args[1], "testSubSet")
		assert.Equal(t, paredQuery.Args[2], "hello")
		assert.Equal(t, paredQuery.Args[3], "1700842078")
	})

	t.Run("empty query", func(t *testing.T) {
		query := ""
		paredQuery := ParseQuery(query)

		assert.Equal(t, "", paredQuery.Command)
		assert.Equal(t, 0, len(paredQuery.Args))
	})
}
