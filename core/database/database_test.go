package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDataBase(t *testing.T) {
	db := Init("../../config/config.yaml")

	t.Run("addSetTest", func(t *testing.T) {
		result := db.AddSet("testSet")

		assert.Equal(t, 1, len(db.Sets))
		assert.Equal(t, "DONE", result)
	})

	t.Run("AddSubSetTest", func(t *testing.T) {
		db.AddSet("testSet")
		result := db.AddSubSet("testSet", "testSubSet")

		assert.Equal(t, 0, len(db.Sets["testSet"]["testSubSet"]))
		assert.Equal(t, "DONE", result)

		result = db.AddSubSet("testInvalidSet", "testSubSet")

		assert.Equal(t, "SETNF", result)
	})

	t.Run("pushElementTest", func(t *testing.T) {
		db.AddSet("testSet")
		db.AddSubSet("testSet", "testSubSet")

		time := time.Now()
		result := db.PushElement("testSet", "testSubSet", Element{value: []byte("testValue"), time: time})

		assert.Equal(t, "DONE", result)
		assert.Equal(t, 1, len(db.Sets["testSet"]["testSubSet"]))
		assert.Equal(t, []byte("testValue"), db.Sets["testSet"]["testSubSet"][0].value)
		assert.Equal(t, time, db.Sets["testSet"]["testSubSet"][0].time)

		result = db.PushElement("invalidTestSet", "invalidTestSubSet", Element{value: []byte("testValue"), time: time})

		assert.Equal(t, "SUBSETNF", result)
		assert.Equal(t, 1, len(db.Sets["testSet"]["testSubSet"]))
		assert.Equal(t, []byte("testValue"), db.Sets["testSet"]["testSubSet"][0].value)
		assert.Equal(t, time, db.Sets["testSet"]["testSubSet"][0].time)
	})

	t.Run("dropSetTest", func(t *testing.T) {
		db.AddSet("testSet")
		db.AddSet("secondTestSet")
		db.AddSet("thirdTestSet")

		result := db.DropSet("testSet")

		assert.Equal(t, 2, len(db.Sets))
		assert.Equal(t, "DONE", result)

		result = db.DropSet("inavlidTestSet")

		assert.Equal(t, "SETNF", result)
		assert.Equal(t, 2, len(db.Sets))
	})

	t.Run("dropSubSetTest", func(t *testing.T) {
		db.AddSet("testSet")
		db.AddSet("secondTestSet")

		db.AddSubSet("testSet", "subSetOne")
		db.AddSubSet("testSet", "subSetTwo")

		result := db.DropSubSet("testSet", "subSetOne")

		assert.Equal(t, "DONE", result)
		assert.Equal(t, 1, len(db.Sets["testSet"]))
		assert.Nil(t, db.Sets["testSet"]["subSetOne"])

		result = db.DropSubSet("secondTestSet", "subSetOne")

		assert.Equal(t, "SUBSETNF", result)
	})

	t.Run("cleanTest", func(t *testing.T) {
		db.AddSet("testSet")
		db.AddSet("secondTestSet")
		db.AddSet("thirdTestSet")

		db.AddSubSet("testSet", "subSetOne")
		db.AddSubSet("testSet", "subSetTwo")

		db.AddSubSet("secondTestSet", "subSetOne")
		db.AddSubSet("secondTestSet", "subSetTwo")

		time := time.Now()
		db.PushElement("testSet", "subSetOne", Element{value: []byte("testValue"), time: time})
		db.PushElement("testSet", "subSetTwo", Element{value: []byte("testValue"), time: time})

		db.PushElement("secondTestSet", "subSetTwo", Element{value: []byte("testValue"), time: time})

		result := db.CleanSubSet("secondTestSet", "subSetTwo")

		assert.Equal(t, "DONE", result)
		assert.Equal(t, 0, len(db.Sets["secondTestSet"]["subSetTwo"]))

		result = db.CleanSet("testSet")

		assert.Equal(t, "DONE", result)
		assert.Equal(t, 0, len(db.Sets["testSet"]))

		result = db.CleanSets()

		assert.Equal(t, "DONE", result)
		assert.Equal(t, 0, len(db.Sets))

		result = db.CleanSet("invalidSet")
		assert.Equal(t, "SETNF", result)

		result = db.CleanSubSet("invalidSet", "invalidSubSet")
		assert.Equal(t, "SUBSETNF", result)
	})
}
