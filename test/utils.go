package test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func AssertInDeltaVec3(t *testing.T, expected core.Vec3, result core.Vec3, delta float32) {
	assert.True(t, expected.InDelta(result, delta), "expected %v, got %v, tolerance %v", expected, result, delta)
}

// In [low, high)
func AssertInSemiInternal(t *testing.T, value, low, high core.Real) {
	assert.GreaterOrEqual(t, value, low)
	assert.Less(t, value, high)
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
