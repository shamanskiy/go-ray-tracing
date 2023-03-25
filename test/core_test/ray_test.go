package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func TestRay_ShouldEvaluate(t *testing.T) {
	ray := core.NewRay(core.NewVec3(1, 2, 3), core.NewVec3(2, 3, 4))

	point := ray.Eval(2)

	assert.Equal(t, core.NewVec3(5, 8, 11), point)
}
