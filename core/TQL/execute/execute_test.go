package execute

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zurvan-lab/TimeTrace/config"
	"github.com/zurvan-lab/TimeTrace/core/TQL/parser"
	"github.com/zurvan-lab/TimeTrace/core/database"
)

func TestExecute(t *testing.T) {
	db := database.Init(config.DefaultConfig())

	q := core.ParseQuery("SET testSet")
	eResult := Execute(q, db)

	_, ok := db.SetsMap()["testSet"]

	assert.Equal(t, "DONE", eResult)
	assert.True(t, ok)

	q2 := core.ParseQuery("CNTS")
	eResult2 := Execute(q2, db)

	assert.Equal(t, "1", eResult2)
}
