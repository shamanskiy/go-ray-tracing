package geometries_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/stretchr/testify/assert"
)

var PLANE_ORIGIN = core.NewVec3(0, 0, 0)
var PLANE_NORMAL = core.NewVec3(0, 1, 0)

func TestPlane_ShouldReturnHit_IfRayIntersectsPlaneWithinParamInterval(t *testing.T) {
	plane := xzPlane()
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, -1, 0))

	hit := plane.TestRay(ray, core.NewInterval(0, 10))

	assert.EqualValues(t, 2., hit.Value().Param)
}

func TestPlane_ShouldReturnNoHit_IfRayIntersectsPlaneOutsideOfParamInterval(t *testing.T) {
	plane := xzPlane()
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, -1, 0))

	hit := plane.TestRay(ray, core.NewInterval(0, 1))

	assert.True(t, hit.Empty())
}

func TestPlane_ShouldEvaluateHitWithPlaneNormal_IfRayHitsPositiveSide(t *testing.T) {
	plane := xzPlane()
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, -1, 0))

	hitRecord := plane.EvaluateHit(ray, 2)

	assert.Equal(t, core.NewVec3(0, 0, 0), hitRecord.Point)
	assert.Equal(t, PLANE_NORMAL, hitRecord.Normal)
}

func TestPlane_ShouldEvaluateHitWithNegatedNormal_IfRayHitsNegativeSide(t *testing.T) {
	plane := xzPlane()
	ray := core.NewRay(core.NewVec3(-2, -2, 0), core.NewVec3(1, 1, 0))

	hitRecord := plane.EvaluateHit(ray, 2)

	assert.Equal(t, core.NewVec3(0, 0, 0), hitRecord.Point)
	assert.Equal(t, PLANE_NORMAL.Mul(-1), hitRecord.Normal)
}

func TestPlane_ShouldReturnNoHit_IfRayIsParallelToPlane(t *testing.T) {
	plane := xzPlane()
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, 0, 0))

	hit := plane.TestRay(ray, core.NewInterval(0, core.Inf()))

	assert.True(t, hit.Empty())
}

func TestPlane_ShouldReturnInfiniteBoundingBox(t *testing.T) {
	plane := xzPlane()

	bbox := plane.BoundingBox()

	assert.Equal(t, core.NewInfiniteBox(), bbox)
}

func xzPlane() geometries.Plane {
	return geometries.NewPlane(PLANE_ORIGIN, PLANE_NORMAL)
}
