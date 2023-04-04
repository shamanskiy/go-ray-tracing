package objects_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/geometries"
	"github.com/stretchr/testify/assert"
)

func TestPlane_ShouldReturnOneHit_IfRayIntersectsPlane(t *testing.T) {
	plane := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, -1, 0))

	hits := plane.TestRay(ray)

	assert.Equal(t, []core.Real{2}, hits)
}

func TestPlane_ShouldEvaluateHitWithPlaneNormal_IfRayHitsPositiveSide(t *testing.T) {
	plane := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, -1, 0))

	hitRecord := plane.EvaluateHit(ray, 2)

	assert.Equal(t, core.NewVec3(0, 0, 0), hitRecord.Point)
	assert.Equal(t, plane.Normal(), hitRecord.Normal)
}

func TestPlane_ShouldEvaluateHitWithNegatedNormal_IfRayHitsNegativeSide(t *testing.T) {
	plane := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	ray := core.NewRay(core.NewVec3(-2, -2, 0), core.NewVec3(1, 1, 0))

	hitRecord := plane.EvaluateHit(ray, 2)

	assert.Equal(t, core.NewVec3(0, 0, 0), hitRecord.Point)
	assert.Equal(t, plane.Normal().Mul(-1), hitRecord.Normal)
}

func TestPlane_ShouldReturnNoHits_IfRayIsParallelToPlane(t *testing.T) {
	plane := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, 0, 0))

	hits := plane.TestRay(ray)

	assert.Empty(t, hits)
}

func TestPlane_ShouldReturnNoHits_IfRayPointsAwayFromPlane(t *testing.T) {
	plane := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	ray := core.NewRay(core.NewVec3(-2, 2, 0), core.NewVec3(1, 1, 0))

	hits := plane.TestRay(ray)

	assert.Empty(t, hits)
}
