package log

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	var buf bytes.Buffer
	log.Logger = log.Output(&buf)

	Trace("a", "a", 4)
	Info("b", nil)
	Info("b", "a", nil)
	Info("c", "b", []byte{1, 2, 3})
	Warn("d", "x")
	Error("error", "A", 3)

	out := buf.String()

	fmt.Println(out)
	assert.Contains(t, out, "010203")
	assert.Contains(t, out, "!INVALID-KEY!")
	assert.Contains(t, out, "!MISSING-VALUE!")
	assert.Contains(t, out, "null")
	assert.NotContains(t, out, "debug")
	assert.Contains(t, out, "info")
	assert.Contains(t, out, "warn")
	assert.Contains(t, out, "error")
}
