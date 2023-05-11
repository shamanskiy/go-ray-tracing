package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func TestInterval(t *testing.T) {
	interval := core.NewInterval(0, 1)

	assert.True(t, interval.Contains(0.5))
	assert.True(t, interval.ContainsStrictly(0.5))
	assert.False(t, interval.Contains(2))
	assert.False(t, interval.ContainsStrictly(0))
	assert.False(t, interval.ContainsStrictly(1))
}
