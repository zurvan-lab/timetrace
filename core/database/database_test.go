package database

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zurvan-lab/TimeTrace/config"
)

func TestDataBase(t *testing.T) {
	db := Init(config.DefaultConfig())

	t.Run("addSetTest", func(t *testing.T) {
		result := db.AddSet([]string{"testSet"})

		assert.Equal(t, 1, len(db.SetsMap()))
		assert.Equal(t, DONE, result)
	})

	t.Run("AddSubSetTest", func(t *testing.T) {
		db.AddSet([]string{"testSet"})
		result := db.AddSubSet([]string{"testSet", "testSubSet"})

		assert.Equal(t, 0, len(db.SetsMap()["testSet"]["testSubSet"]))
		assert.Equal(t, DONE, result)

		result = db.AddSubSet([]string{"testInvalidSet", "testSubSet"})

		assert.Equal(t, SET_NOT_FOUND, result)
	})

	t.Run("pushElementTest", func(t *testing.T) {
		db.AddSet([]string{"testSet"})
		db.AddSubSet([]string{"testSet", "testSubSet"})

		timeStr := fmt.Sprintf("%d", time.Now().Unix())
		result := db.PushElement([]string{"testSet", "testSubSet", "testValue", timeStr})

		assert.Equal(t, DONE, result)
		assert.Equal(t, 1, len(db.SetsMap()["testSet"]["testSubSet"]))
		assert.Equal(t, "testValue", db.SetsMap()["testSet"]["testSubSet"][0].value)

		elementTime := strconv.Itoa(int(db.SetsMap()["testSet"]["testSubSet"][0].time.Unix()))
		assert.Equal(t, timeStr, elementTime)

		result = db.PushElement([]string{"invalidTestSet", "invalidTestSubSet", "testValue", timeStr})

		assert.Equal(t, SUB_SET_NOT_FOUND, result)
		assert.Equal(t, 1, len(db.SetsMap()["testSet"]["testSubSet"]))
		assert.Equal(t, "testValue", db.SetsMap()["testSet"]["testSubSet"][0].value)

		elementTime = strconv.Itoa(int(db.SetsMap()["testSet"]["testSubSet"][0].time.Unix()))
		assert.Equal(t, timeStr, elementTime)
	})

	t.Run("dropSetTest", func(t *testing.T) {
		db.AddSet([]string{"testSet"})
		db.AddSet([]string{"secondTestSet"})
		db.AddSet([]string{"thirdTestSet"})

		result := db.DropSet([]string{"testSet"})

		assert.Equal(t, 2, len(db.SetsMap()))
		assert.Equal(t, DONE, result)

		result = db.DropSet([]string{"inavlidTestSet"})

		assert.Equal(t, SET_NOT_FOUND, result)
		assert.Equal(t, 2, len(db.SetsMap()))
	})

	t.Run("dropSubSetTest", func(t *testing.T) {
		db.AddSet([]string{"testSet"})
		db.AddSet([]string{"secondTestSet"})

		db.AddSubSet([]string{"testSet", "subSetOne"})
		db.AddSubSet([]string{"testSet", "subSetTwo"})

		result := db.DropSubSet([]string{"testSet", "subSetOne"})

		assert.Equal(t, DONE, result)
		assert.Equal(t, 1, len(db.SetsMap()["testSet"]))
		assert.Nil(t, db.SetsMap()["testSet"]["subSetOne"])

		result = db.DropSubSet([]string{"secondTestSet", "subSetOne"})

		assert.Equal(t, SUB_SET_NOT_FOUND, result)
	})

	t.Run("cleanTest", func(t *testing.T) {
		db.AddSet([]string{"testSet"})
		db.AddSet([]string{"secondTestSet"})
		db.AddSet([]string{"thirdTestSet"})

		db.AddSubSet([]string{"testSet", "subSetOne"})
		db.AddSubSet([]string{"testSet", "subSetTwo"})

		db.AddSubSet([]string{"secondTestSet", "subSetOne"})
		db.AddSubSet([]string{"secondTestSet", "subSetTwo"})

		time := time.Now()
		timeStr := strconv.Itoa(int(time.Unix()))
		db.PushElement([]string{"testSet", "subSetOne", "testValue", timeStr})
		db.PushElement([]string{"testSet", "subSetTwo", "testValue", timeStr})

		db.PushElement([]string{"secondTestSet", "subSetTwo", "testValue", timeStr})

		result := db.CleanSubSet([]string{"secondTestSet", "subSetTwo"})

		assert.Equal(t, DONE, result)
		assert.Equal(t, 0, len(db.SetsMap()["secondTestSet"]["subSetTwo"]))

		result = db.CleanSet([]string{"testSet"})

		assert.Equal(t, DONE, result)
		assert.Equal(t, 0, len(db.SetsMap()["testSet"]))

		result = db.CleanSets([]string{})

		assert.Equal(t, DONE, result)
		assert.Equal(t, 0, len(db.SetsMap()))

		result = db.CleanSet([]string{"invalidSet"})
		assert.Equal(t, SET_NOT_FOUND, result)

		result = db.CleanSubSet([]string{"invalidSet", "invalidSubSet"})
		assert.Equal(t, SUB_SET_NOT_FOUND, result)
	})

	t.Run("countTest", func(t *testing.T) {
		db.AddSet([]string{"testSet"})
		db.AddSet([]string{"secondTestSet"})
		db.AddSet([]string{"thirdTestSet"})

		db.AddSubSet([]string{"testSet", "subSetOne"})
		db.AddSubSet([]string{"testSet", "subSetTwo"})

		db.AddSubSet([]string{"secondTestSet", "subSetOne"})
		db.AddSubSet([]string{"secondTestSet", "subSetTwo"})

		time := time.Now()
		timeStr := strconv.Itoa(int(time.Unix()))
		db.PushElement([]string{"testSet", "subSetOne", "testValue", timeStr})
		db.PushElement([]string{"testSet", "subSetTwo", "testValue", timeStr})

		db.PushElement([]string{"secondTestSet", "subSetTwo", "testValue", timeStr})

		result := db.CountSets([]string{})
		assert.Equal(t, "3", result)

		result = db.CountSubSets([]string{"testSet"})
		assert.Equal(t, "2", result)

		result = db.CountElements([]string{"testSet", "subSetOne"})
		assert.Equal(t, "1", result)
	})

	t.Run("getTest", func(t *testing.T) {
		db.AddSet([]string{"testSet"})
		db.AddSubSet([]string{"testSet", "subSetOne"})

		time := time.Now()
		timeStr := strconv.Itoa(int(time.Unix()))
		for i := 0; i < 50; i++ {
			db.PushElement([]string{"testSet", "subSetOne", "testValue", timeStr})
		}

		result := db.GetElements([]string{"testSet", "subSetOne"})

		trimmedResult, _ := strings.CutPrefix(result, "  ")
		trimmedResult, _ = strings.CutPrefix(trimmedResult, "  ")
		fmt.Print(trimmedResult)
		resultsArray := strings.Split(trimmedResult, "  ")

		assert.NotEqual(t, INVALID, result)
		assert.NotEqual(t, SUB_SET_NOT_FOUND, result)
		assert.Equal(t, 50, len(resultsArray))

		result = db.GetElements([]string{"testSet", "subSetOne", "10"})

		trimmedResult, _ = strings.CutPrefix(result, "  ")
		trimmedResult, _ = strings.CutPrefix(trimmedResult, "  ")
		resultsArray = strings.Split(trimmedResult, "  ")

		assert.NotEqual(t, INVALID, result)
		assert.NotEqual(t, SUB_SET_NOT_FOUND, result)
		assert.Equal(t, 10, len(resultsArray))
	})
}
