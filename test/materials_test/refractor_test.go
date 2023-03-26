package materials_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

const (
	GLASS_REFRACTION_INDEX = 1.5
	AIR_REFRACTION_INDEX   = 1.
)

var normalUp = core.NewVec3(0, 1, 0)

func TestRefractionCalculator_ShouldPushRayToNormal_WhenRayEntersGlass(t *testing.T) {
	refractor := materials.NewRefractionCalculator(GLASS_REFRACTION_INDEX)
	incidentVector := core.NewVec3(1, -1, 0)

	refraction := refractor.Refract(incidentVector, normalUp)

	test.AssertInDeltaVec3(t, core.NewVec3(0.666666, -1.247219, 0), *refraction.Direction, core.Tolerance)
	assert.InDelta(t, 0.04207, refraction.ReflectionRatio, core.Tolerance)
}

func TestRefractionCalculator_ShouldPushRayFromNormal_WhenRayExitsGlass(t *testing.T) {
	refractor := materials.NewRefractionCalculator(GLASS_REFRACTION_INDEX)
	incidentVector := core.NewVec3(3, 4, 0.)

	refraction := refractor.Refract(incidentVector, normalUp)

	test.AssertInDeltaVec3(t, core.NewVec3(4.5, 2.179448, 0), *refraction.Direction, core.Tolerance)
	assert.InDelta(t, 0.0948391, refraction.ReflectionRatio, core.Tolerance)
}

func TestRefractionCalculator_ShouldReturnFullInternalReflection_WhenRayExitsGlassAndIncidentAngleTooBig(t *testing.T) {
	refractor := materials.NewRefractionCalculator(GLASS_REFRACTION_INDEX)
	incidentVector := core.NewVec3(1., 1., 0.)

	refraction := refractor.Refract(incidentVector, normalUp)

	assert.True(t, refraction.FullInternalReflection())
}

func TestRefractionCalculator_ShouldNotChangeRayDirection_WhenRayEntersAir(t *testing.T) {
	refractor := materials.NewRefractionCalculator(AIR_REFRACTION_INDEX)
	incidentVector := core.NewVec3(1., -1., 0.)

	refraction := refractor.Refract(incidentVector, normalUp)

	test.AssertInDeltaVec3(t, incidentVector, *refraction.Direction, core.Tolerance)
	assert.InDelta(t, 0.002155, refraction.ReflectionRatio, core.Tolerance)
}

func TestRefractionCalculator_ShouldNotChangeRayDirection_WhenRayExitsAir(t *testing.T) {
	refractor := materials.NewRefractionCalculator(AIR_REFRACTION_INDEX)
	incidentVector := core.NewVec3(3, 4, 0.)

	refraction := refractor.Refract(incidentVector, normalUp)

	test.AssertInDeltaVec3(t, incidentVector, *refraction.Direction, core.Tolerance)
	assert.InDelta(t, 0.00032, refraction.ReflectionRatio, core.Tolerance)
}
