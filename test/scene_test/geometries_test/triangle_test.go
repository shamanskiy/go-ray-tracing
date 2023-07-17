package geometries

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/stretchr/testify/assert"
)

func TestTriangleBoundingBox(t *testing.T) {
	triangle := xyzTriangle()

	bbox := triangle.BoundingBox()

	expectedBBox := core.NewBox(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))
	assert.Equal(t, expectedBBox, bbox)
}

func TestTriangle_RayShouldMiss(t *testing.T) {
	triangle := xyzTriangle()
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(-1, 0, 0))

	hit := triangle.TestRay(ray, core.NewInterval(0, 10))

	assert.True(t, hit.Empty())
}

func TestTriangle_ShouldReturnNoHitBecauseOutsideOfParamInterval(t *testing.T) {
	triangle := xyzTriangle()
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))

	hit := triangle.TestRay(ray, core.NewInterval(0, 0.1))

	assert.True(t, hit.Empty())
}

func TestTriangle_ShouldReturnHit(t *testing.T) {
	triangle := xyzTriangle()
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))

	hit := triangle.TestRay(ray, core.NewInterval(0, 10))

	assert.InDelta(t, 1./3, hit.Value().Param, core.Tolerance)
	assert.Equal(t, core.NewVec3(1./3, 1./3, 1./3), hit.Value().Point)
	normValue := core.Sqrt(3) / 3
	assert.Equal(t, core.NewVec3(normValue, normValue, normValue), hit.Value().Normal)
}

func xyzTriangle() geometries.Triangle {
	return geometries.NewTriangleWithNormals(
		core.NewVec3(1, 0, 0),
		core.NewVec3(0, 1, 0),
		core.NewVec3(0, 0, 1),
		core.NewVec3(1, 1, 1),
		core.NewVec3(1, 1, 1),
		core.NewVec3(1, 1, 1))
}
