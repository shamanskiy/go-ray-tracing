package materials_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/stretchr/testify/assert"
)

var anyDirection = core.NewVec3(10, 20, 30)
var anyColor = color.Red
var hitPoint = core.NewVec3(0, 1, 2)
var normalAtHitPointUp = core.NewVec3(0, 1, 0)

func TestDiffusive_ShouldReflectRayInNormalDirection_WhenNotRandom(t *testing.T) {
	material := materials.NewDiffusive(anyColor, random.NewFakeRandomGenerator())

	reflection := material.Reflect(anyDirection, hitPoint, normalAtHitPointUp)

	expected := materials.Reflection{
		Ray:         core.NewRay(hitPoint, normalAtHitPointUp),
		Attenuation: material.Color(),
	}
	assert.Equal(t, expected, *reflection)
}

func TestDiffusive_ShouldReflectRayWithinUnitSphereOfNormal_WhenRandom(t *testing.T) {
	material := materials.NewDiffusive(anyColor, random.NewRandomGenerator())

	reflection := material.Reflect(anyDirection, hitPoint, normalAtHitPointUp)

	randomPerturbation := reflection.Ray.Direction().Sub(normalAtHitPointUp).Len()
	assert.Less(t, randomPerturbation, core.Real(1))
	assert.Equal(t, material.Color(), reflection.Attenuation)
	assert.Equal(t, hitPoint, reflection.Ray.Origin())
}
