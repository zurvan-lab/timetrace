package execute

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zurvan-lab/TimeTrace/core/TQL/parser"
	"github.com/zurvan-lab/TimeTrace/core/database"
)

func TestExecute(t *testing.T) {
	db := database.Init("../../../config/config.yaml")

	query := "SET testSet"
	q := core.ParseQuery(query)

	eResult := Execute(q, db)
	_, ok := db.SetsMap()["testSet"]

	assert.Equal(t, "DONE", eResult)
	assert.True(t, ok)
}
