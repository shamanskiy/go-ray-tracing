package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func TestRayShouldHitBox(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))
	box := core.NewBox(core.NewVec3(1, 1, 1), core.NewVec3(3, 3, 3))

	assert.True(t, ray.Hits(box, 0, 10))
}

func TestRayShouldNotHitBox_IfHitParamOutsideOfRequestedInterval(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 1, 1))
	box := core.NewBox(core.NewVec3(1, 1, 1), core.NewVec3(3, 3, 3))

	assert.False(t, ray.Hits(box, 0, 0.5))
}

func TestRayShouldNotHitBox_IfRayPointsInWrongDirection(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(-1, 1, 1))
	box := core.NewBox(core.NewVec3(1, 1, 1), core.NewVec3(3, 3, 3))

	assert.False(t, ray.Hits(box, 0, 10))
}

func TestRayShouldHitBox_IfRayParallelToBoxSide(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, 0, 10))
}

func TestRayShouldHitBox_IfRayTravelsInNegativeDirection(t *testing.T) {
	ray := core.NewRay(core.NewVec3(5, 0, 0), core.NewVec3(-1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, 0, 10))
}

func TestRayShouldHitBox_IfRayOriginIsOnBoxBoundary(t *testing.T) {
	ray := core.NewRay(core.NewVec3(1, 0, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, 0, 10))
}

func TestRayShouldHitBox_IfRayGoesThroughBoxSide(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 1, 0), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, 0, 10))
}

func TestRayShouldHitBox_IfRayGoesThroughBoxEdge(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 1, 1), core.NewVec3(1, 0, 0))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.True(t, ray.Hits(box, 0, 10))
}

// Why ray-box hit is detected for sides and edges but not for corners? No idea
func TestRayShouldNotHitBox_IfRayGoesThroughBoxCorner(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, -1, -1))
	box := core.NewBox(core.NewVec3(1, -1, -1), core.NewVec3(3, 1, 1))

	assert.False(t, ray.Hits(box, 0, 10))
}
