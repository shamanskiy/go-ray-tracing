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
