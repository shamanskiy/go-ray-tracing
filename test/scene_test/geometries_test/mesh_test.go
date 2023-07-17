package geometries_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/stretchr/testify/assert"
)

func TestMeshQuad_ShouldComputeBBox(t *testing.T) {
	quad := unitSquareXYQuad()

	expectedBBox := core.NewBox(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 0))
	assert.Equal(t, expectedBBox, quad.BoundingBox())
}

func TestMeshQuad_ShouldTestRay(t *testing.T) {
	quad := unitSquareXYQuad()
	ray := core.NewRay(core.NewVec3(0.5, 0.5, 2), core.NewVec3(0, 0, -1))

	hit := quad.TestRay(ray, core.NewInterval(0, core.Inf()))

	assert.EqualValues(t, 2, hit.Value().Param)
	assert.Equal(t, core.NewVec3(0.5, 0.5, 0), hit.Value().Point)
	assert.Equal(t, core.NewVec3(0, 0, 1), hit.Value().Normal)
}

func unitSquareXYQuad() geometries.Mesh {
	return geometries.NewQuad(
		core.NewVec3(0, 0, 0),
		core.NewVec3(1, 0, 0),
		core.NewVec3(1, 1, 0),
		core.NewVec3(0, 1, 0))
}
