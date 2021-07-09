package data

import (
	"testing"

	"gotest.tools/assert"
)

func TestValid(t *testing.T) {
	var original, transformed Figure
	var epsilon = 15000

	assert.Assert(t, transformed.IsValid(original, epsilon))
}
