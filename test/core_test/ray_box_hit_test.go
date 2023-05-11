package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func TestRayShouldHitBox(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))
	box := core.NewBox(core.NewVec3(1, 1, 1), core.NewVec3(3, 3, 3))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldNotHitBox_IfHitParamOutsideOfRequestedInterval(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))
	box := core.NewBox(core.NewVec3(1, 1, 1), core.NewVec3(3, 3, 3))

	assert.False(t, ray.Hits(box, core.NewInterval(0, 0.5)))
}

func TestRayShouldNotHitBox_IfRayPointsInWrongDirection(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(-1, 1, 1))
	box := core.NewBox(core.NewVec3(1, 1, 1), core.NewVec3(3, 3, 3))

	assert.False(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitBox_IfRayParallelToBoxSide(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitBox_IfRayTravelsInNegativeDirection(t *testing.T) {
	ray := core.NewRay(core.NewVec3(5, 0, 0), core.NewVec3(-1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitBox_IfRayOriginIsOnBoxBoundary(t *testing.T) {
	ray := core.NewRay(core.NewVec3(1, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitBox_IfRayGoesThroughBoxSide(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 1, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitBox_IfRayGoesThroughBoxEdge(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 1, 1), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldNotHitBox_IfRayGoesThroughBoxCorner(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, -1, -1))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitDegenerateFlatBox(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(1, 1, 1))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitDegenerateLineBox(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, 0, 0), core.NewVec3(3, 0, 0))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitDegeneratePointBox(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, 0, 0), core.NewVec3(1, 0, 0))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitDegenerateFlatBox_IfRayIsNotNormalToBoxPlane(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, -1, -1), core.NewVec3(1, 1, 1))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(1, 1, 1))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 10)))
}

func TestRayShouldHitInfiniteBox(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(-core.Inf(), -core.Inf(), -core.Inf()), core.NewVec3(core.Inf(), core.Inf(), core.Inf()))

	assert.True(t, ray.Hits(box, core.NewInterval(0, 1)))
}

func TestInfiniteRayShouldHitInfiniteBox(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(-core.Inf(), -core.Inf(), -core.Inf()), core.NewVec3(core.Inf(), core.Inf(), core.Inf()))

	assert.True(t, ray.Hits(box, core.NewInterval(0, core.Inf())))
}
