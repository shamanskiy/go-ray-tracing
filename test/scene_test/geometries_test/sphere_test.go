package geometries_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/stretchr/testify/assert"
)

func TestSphere_ShouldReturnClosestHit_IfHasTwoIntersectionsWithinParamInterval(t *testing.T) {
	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 2)
	ray := core.NewRay(core.NewVec3(4, 0, 0), core.NewVec3(-1, 0, 0))

	hit := sphere.TestRay(ray, core.NewInterval(0, 10))

	assert.True(t, hit.HasHit)
	assert.EqualValues(t, 2., hit.Param)
}

func TestSphere_ShouldReturnFarthestHit_IfClosestHitOutsideParamInternal(t *testing.T) {
	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 2)
	ray := core.NewRay(core.NewVec3(4, 0, 0), core.NewVec3(-1, 0, 0))

	hit := sphere.TestRay(ray, core.NewInterval(3, 10))

	assert.True(t, hit.HasHit)
	assert.EqualValues(t, 6., hit.Param)
}

func TestSphere_ShouldReturnNoHit_IfBothHitsOutsideOfParamInterval(t *testing.T) {
	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 2)
	ray := core.NewRay(core.NewVec3(4, 0, 0), core.NewVec3(-1, 0, 0))

	hit := sphere.TestRay(ray, core.NewInterval(0, 1))

	assert.False(t, hit.HasHit)
}

func TestSphere_ShouldReturnHit_IfRayTouchesSphere(t *testing.T) {
	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 2)
	ray := core.NewRay(core.NewVec3(4, 2, 0), core.NewVec3(-1, 0, 0))

	hit := sphere.TestRay(ray, core.NewInterval(0, 10))

	assert.True(t, hit.HasHit)
	assert.EqualValues(t, 4., hit.Param)
}

func TestSphere_ShouldReturnNoHits_IfRayDoesNotIntersectSphere(t *testing.T) {
	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 2)
	ray := core.NewRay(core.NewVec3(4, 4, 0), core.NewVec3(-1, 0, 0))

	hit := sphere.TestRay(ray, core.NewInterval(0, 10))

	assert.False(t, hit.HasHit)
}

func TestSphere_ShouldReturnNoHits_IfRayPointsDirectlyAgainstSphere(t *testing.T) {
	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 2)
	ray := core.NewRay(core.NewVec3(4, 0, 0), core.NewVec3(1, 0, 0))

	hit := sphere.TestRay(ray, core.NewInterval(0, 10))

	assert.False(t, hit.HasHit)
}

func TestSphere_ShouldEvaluateHit(t *testing.T) {
	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 2)
	ray := core.NewRay(core.NewVec3(4, 0, 0), core.NewVec3(-1, 0, 0))

	hitRecord := sphere.EvaluateHit(ray, 2)

	assert.Equal(t, core.NewVec3(2, 0, 0), hitRecord.Point)
	assert.Equal(t, core.NewVec3(1, 0, 0), hitRecord.Normal)
}

func TestSphere_ShouldComputeBoundingBox(t *testing.T) {
	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 2)

	bbox := sphere.BoundingBox()

	expectedBBox := core.NewBox(core.NewVec3(-2, -2, -2), core.NewVec3(2, 2, 2))
	assert.Equal(t, expectedBBox, bbox)
}
