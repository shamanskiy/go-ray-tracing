package objects_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/stretchr/testify/assert"
)

func TestPlane_ShouldReturnOneHit(t *testing.T) {
	plane := objects.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, -1, 0))

	hits := plane.TestRay(ray)

	assert.Equal(t, []core.Real{2}, hits)
}
