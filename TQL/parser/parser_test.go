package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQuery(t *testing.T) {
	query := "PUSH testSet testSubSet hello NOW"
	paredQuery := ParseQuery(query)

	assert.Equal(t, paredQuery.Command, "PUSH")
	assert.Equal(t, paredQuery.Args[0], "testSet")
	assert.Equal(t, paredQuery.Args[1], "testSubSet")
	assert.Equal(t, paredQuery.Args[2], "hello")
	assert.Equal(t, paredQuery.Args[3], "NOW")
}
