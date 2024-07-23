package execute

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zurvan-lab/timetrace/config"
	"github.com/zurvan-lab/timetrace/core/TQL/parser"
	"github.com/zurvan-lab/timetrace/core/database"
)

func TestExecute(t *testing.T) {
	db := database.Init(config.DefaultConfig())

	q := core.ParseQuery("SET testSet")
	eResult := Execute(q, db)

	_, ok := db.SetsMap()["testSet"]

	assert.Equal(t, "OK", eResult)
	assert.True(t, ok)

	q2 := core.ParseQuery("CNTS")
	eResult2 := Execute(q2, db)

	assert.Equal(t, "1", eResult2)
}
