package geometries

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/stretchr/testify/assert"
)

func TestTriangleBoundingBox(t *testing.T) {
	triangle := geometries.NewTriangle(
		core.NewVec3(1, 0, 0),
		core.NewVec3(0, 1, 0),
		core.NewVec3(0, 0, 1))

	bbox := triangle.BoundingBox()

	expectedBBox := core.NewBox(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))
	assert.Equal(t, expectedBBox, bbox)
}

func TestTriangle_RayShouldMiss(t *testing.T) {
	triangle := geometries.NewTriangle(
		core.NewVec3(1, 0, 0),
		core.NewVec3(0, 1, 0),
		core.NewVec3(0, 0, 1))
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(-1, 0, 0))

	hit := triangle.TestRay(ray, core.NewInterval(0, 10))

	assert.True(t, hit.Empty())
}

func TestTriangle_ShouldReturnNoHitBecauseOutsideOfParamInterval(t *testing.T) {
	triangle := geometries.NewTriangle(
		core.NewVec3(1, 0, 0),
		core.NewVec3(0, 1, 0),
		core.NewVec3(0, 0, 1))
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))

	hit := triangle.TestRay(ray, core.NewInterval(0, 0.1))

	assert.True(t, hit.Empty())
}

// func TestTriangle_ShouldReturnHit(t *testing.T) {
// 	triangle := geometries.NewTriangle(
// 		core.NewVec3(1, 0, 0),
// 		core.NewVec3(0, 1, 0),
// 		core.NewVec3(0, 0, 1))
// 	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))

// 	hit := triangle.TestRay(ray, core.NewInterval(0, 10))

// 	assert.Equal(t, math32.Sqrt(2), hit.Value().Param)
// }
