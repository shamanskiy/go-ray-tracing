package test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func AssertInDeltaVec3(t *testing.T, expected core.Vec3, result core.Vec3, delta float32) {
	assert.True(t, expected.InDelta(result, delta), "expected %v, got %v, tolerance %v", expected, result, delta)
}
