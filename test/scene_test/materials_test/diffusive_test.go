package materials_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/stretchr/testify/assert"
)

var RAY_DIRECTION = core.NewVec3(10, 20, 30)
var MATERIAL_COLOR = color.Red
var HIT_POINT = core.NewVec3(0, 1, 2)
var NORMAL_AT_HIT_POINT = core.NewVec3(0, 1, 0)

func TestDiffusive_ShouldReflectRayInNormalDirection_WhenNotRandom(t *testing.T) {
	material := materials.NewDiffusive(MATERIAL_COLOR, random.NewFakeRandomGenerator())

	reflection := material.Reflect(RAY_DIRECTION, HIT_POINT, NORMAL_AT_HIT_POINT)

	expected := materials.Reflection{
		Type:  materials.Scattered,
		Ray:   core.NewRay(HIT_POINT, NORMAL_AT_HIT_POINT),
		Color: MATERIAL_COLOR,
	}
	assert.Equal(t, expected, reflection)
}

func TestDiffusive_ShouldReflectRayWithinUnitSphereOfNormal_WhenRandom(t *testing.T) {
	material := materials.NewDiffusive(MATERIAL_COLOR, random.NewRandomGenerator())

	reflection := material.Reflect(RAY_DIRECTION, HIT_POINT, NORMAL_AT_HIT_POINT)

	randomPerturbation := reflection.Ray.Direction().Sub(NORMAL_AT_HIT_POINT).Len()
	assert.Less(t, randomPerturbation, core.Real(1))
	assert.Equal(t, MATERIAL_COLOR, reflection.Color)
	assert.Equal(t, HIT_POINT, reflection.Ray.Origin())
}
