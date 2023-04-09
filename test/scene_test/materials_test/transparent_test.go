package materials_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

var INCIDENT_DIRECTION = core.NewVec3(1.0, -1.0, 0.0)
var REFLECTED_DIRECTION = core.NewVec3(1.0, 1.0, 0.0)
var REFRACTED_DIRECTION = core.NewVec3(0.666666, -1.247219, 0)

func TestTransparent_RefractionIndexCantBeLessThanOne(t *testing.T) {
	assert.Panics(t, func() {
		materials.NewTransparent(0.5, MATERIAL_COLOR, random.NewRandomGenerator())
	})
}

func TestTransparent_ShouldReturnRefractedRay_WhenRandomReturnsOne(t *testing.T) {
	randomizer := random.NewFakeRandomGenerator()
	randomizer.RealValue = 1
	material := materials.NewTransparent(GLASS_REFRACTION_INDEX, MATERIAL_COLOR, randomizer)

	reflection := material.Reflect(INCIDENT_DIRECTION, HIT_POINT, NORMAL_AT_HIT_POINT)

	test.AssertInDeltaVec3(t, REFRACTED_DIRECTION, reflection.Ray.Direction(), core.Tolerance)
	assert.Equal(t, HIT_POINT, reflection.Ray.Origin())
	assert.Equal(t, MATERIAL_COLOR, reflection.Color)
	assert.Equal(t, materials.Scattered, reflection.Type)
}

func TestTransparent_ShouldReturnReflectedRay_WhenRandomReturnsZero(t *testing.T) {
	randomizer := random.NewFakeRandomGenerator()
	randomizer.RealValue = 0
	material := materials.NewTransparent(GLASS_REFRACTION_INDEX, MATERIAL_COLOR, randomizer)

	reflection := material.Reflect(INCIDENT_DIRECTION, HIT_POINT, NORMAL_AT_HIT_POINT)

	assert.Equal(t, REFLECTED_DIRECTION, reflection.Ray.Direction())
	assert.Equal(t, HIT_POINT, reflection.Ray.Origin())
	assert.Equal(t, MATERIAL_COLOR, reflection.Color)
	assert.Equal(t, materials.Scattered, reflection.Type)
}

func TestTransparent_ShouldReturnRefractedOrReflectedRay_WhenRandomEnabled(t *testing.T) {
	material := materials.NewTransparent(GLASS_REFRACTION_INDEX, MATERIAL_COLOR, random.NewRandomGenerator())

	reflection := material.Reflect(INCIDENT_DIRECTION, HIT_POINT, NORMAL_AT_HIT_POINT)

	assert.True(t, reflection.Ray.Direction().InDelta(REFLECTED_DIRECTION, core.Tolerance) ||
		reflection.Ray.Direction().InDelta(REFRACTED_DIRECTION, core.Tolerance))
	assert.Equal(t, HIT_POINT, reflection.Ray.Origin())
	assert.Equal(t, MATERIAL_COLOR, reflection.Color)
	assert.Equal(t, materials.Scattered, reflection.Type)
}

func TestTransparent_ShouldReturnReflectedRay_WhenRayExitsMaterialAtTooLargeAngle(t *testing.T) {
	material := materials.NewTransparent(GLASS_REFRACTION_INDEX, MATERIAL_COLOR, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(1, 1, 0)

	reflection := material.Reflect(incidentDirection, HIT_POINT, NORMAL_AT_HIT_POINT)

	reflectedDirection := core.NewVec3(1, -1, 0)
	assert.Equal(t, reflectedDirection, reflection.Ray.Direction())
	assert.Equal(t, HIT_POINT, reflection.Ray.Origin())
	assert.Equal(t, MATERIAL_COLOR, reflection.Color)
	assert.Equal(t, materials.Scattered, reflection.Type)
}
