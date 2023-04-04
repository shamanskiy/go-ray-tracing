package materials_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

var incidentDirection = core.NewVec3(1.0, -1.0, 0.0)
var reflectedDirection = core.NewVec3(1.0, 1.0, 0.0)
var refractedDirection = core.NewVec3(0.666666, -1.247219, 0)

func TestTransparent_RefractionIndexCantBeLessThanOne(t *testing.T) {
	assert.Panics(t, func() {
		materials.NewTransparent(0.5, anyColor, random.NewRandomGenerator())
	})
}

func TestTransparent_ShouldReturnRefractedRay_WhenRandomReturnsOne(t *testing.T) {
	randomizer := random.NewFakeRandomGenerator()
	randomizer.RealValue = 1
	material := materials.NewTransparent(GLASS_REFRACTION_INDEX, anyColor, randomizer)

	reflection := material.Reflect(incidentDirection, hitPoint, normalAtHitPointUp)

	test.AssertInDeltaVec3(t, refractedDirection, reflection.Ray.Direction(), core.Tolerance)
	assert.Equal(t, hitPoint, reflection.Ray.Origin())
	assert.Equal(t, material.Color(), reflection.Color)
}

func TestTransparent_ShouldReturnReflectedRay_WhenRandomReturnsZero(t *testing.T) {
	randomizer := random.NewFakeRandomGenerator()
	randomizer.RealValue = 0
	material := materials.NewTransparent(GLASS_REFRACTION_INDEX, anyColor, randomizer)

	reflection := material.Reflect(incidentDirection, hitPoint, normalAtHitPointUp)

	assert.Equal(t, reflectedDirection, reflection.Ray.Direction())
	assert.Equal(t, hitPoint, reflection.Ray.Origin())
	assert.Equal(t, material.Color(), reflection.Color)
}

func TestTransparent_ShouldReturnRefractedOrReflectedRay_WhenRandomEnabled(t *testing.T) {
	material := materials.NewTransparent(GLASS_REFRACTION_INDEX, anyColor, random.NewRandomGenerator())

	reflection := material.Reflect(incidentDirection, hitPoint, normalAtHitPointUp)

	assert.True(t, reflection.Ray.Direction().InDelta(reflectedDirection, core.Tolerance) ||
		reflection.Ray.Direction().InDelta(refractedDirection, core.Tolerance))
	assert.Equal(t, hitPoint, reflection.Ray.Origin())
	assert.Equal(t, material.Color(), reflection.Color)
}

func TestTransparent_ShouldReturnReflectedRay_WhenRayExitsMaterialAtTooLargeAngle(t *testing.T) {
	material := materials.NewTransparent(GLASS_REFRACTION_INDEX, anyColor, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(1, 1, 0)

	reflection := material.Reflect(incidentDirection, hitPoint, normalAtHitPointUp)

	reflectedDirection := core.NewVec3(1, -1, 0)
	assert.Equal(t, reflectedDirection, reflection.Ray.Direction())
	assert.Equal(t, hitPoint, reflection.Ray.Origin())
	assert.Equal(t, material.Color(), reflection.Color)
}
