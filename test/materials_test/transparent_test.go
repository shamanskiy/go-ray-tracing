package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

func TestTransparent_RefractionIndexCantBeLessThanOne(t *testing.T) {
	assert.Panics(t, func() {
		materials.NewTransparent(0.5, random.NewRandomGenerator())
	})
}

func TestRefract_OutsideToInside_Glass(t *testing.T) {
	normal := core.NewVec3(0., 1., 0.)
	refractionIndex := core.Real(1.5)

	incidentVector := core.NewVec3(1., -1., 0.)
	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.NewVec3(0.666666, -1.247219, 0), *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.04207, reflectionRatio, core.Tolerance)

	incidentVector = core.NewVec3(1., -0.1, 0.)
	refractedDirection, reflectionRatio = materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.NewVec3(0.666666, -0.752034, 0), *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.60843, reflectionRatio, core.Tolerance)

	incidentVector = core.NewVec3(1., -0.01, 0.)
	refractedDirection, reflectionRatio = materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.NewVec3(0.666666, -0.745423, 0), *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.95295, reflectionRatio, core.Tolerance)
}

func TestRefract_OutsideToInside_Air(t *testing.T) {
	normal := core.NewVec3(0., 1., 0.)
	refractionIndex := core.Real(1.)

	incidentVector := core.NewVec3(1., -1., 0.)

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.NewVec3(1, -1, 0), *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.002155, reflectionRatio, core.Tolerance)
}

func TestRefract_InsideToOutside_Glass(t *testing.T) {
	normal := core.NewVec3(0., 1., 0.)
	refractionIndex := core.Real(1.5)

	incidentVector := core.NewVec3(3, 4, 0.)

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.NewVec3(4.5, 2.179448, 0), *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.0948391, reflectionRatio, core.Tolerance)
}

func TestRefract_InsideToOutside_Air(t *testing.T) {
	normal := core.NewVec3(0., 1., 0.)
	refractionIndex := core.Real(1.)

	incidentVector := core.NewVec3(3, 4, 0.)

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.AssertInDeltaVec3(t, core.NewVec3(3, 4, 0), *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.00032, reflectionRatio, core.Tolerance)
}

func TestRefract_InsideToOutside_Glass_ShouldFullyReflect_AngleTooBig(t *testing.T) {
	normal := core.NewVec3(0., 1., 0.)
	refractionIndex := core.Real(1.5)

	incidentVector := core.NewVec3(1., 1., 0.)

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	assert.Nil(t, refractedDirection)
	assert.InDelta(t, 1, reflectionRatio, core.Tolerance)
}

func TestSchlickLaw(t *testing.T) {
	refractionIndex := core.Real(1.5)

	cosIn := core.Real(1.0)
	reflectionRatio := materials.SchlickLaw(cosIn, refractionIndex)
	assert.InDelta(t, 0.04, reflectionRatio, core.Tolerance)

	cosIn = core.Real(0.0)
	reflectionRatio = materials.SchlickLaw(cosIn, refractionIndex)
	assert.InDelta(t, 1, reflectionRatio, core.Tolerance)

	cosIn = core.Real(0.5)
	reflectionRatio = materials.SchlickLaw(cosIn, refractionIndex)
	assert.InDelta(t, 0.07, reflectionRatio, core.Tolerance)
}

func TestTransparent_RayGetsRefractedOrReflected(t *testing.T) {
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewTransparent(1.5, randomizer)
	incidentDirection := core.NewVec3(1.0, -1.0, 0.0)
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(1.0, 2.0, 3.0),
		Normal: core.NewVec3(0.0, 1.0, 0.0),
	}

	reflection := material.Reflect(incidentDirection, hit.Point, hit.Normal)

	assert.NotNil(t, reflection)
	assert.Equal(t, color.White, reflection.Attenuation)
	assert.Equal(t, hit.Point, reflection.Ray.Origin())

	refractedDirection := core.NewVec3(0.666666, -1.247219, 0)
	reflectedDirection := core.NewVec3(1.0, 1.0, 0.0)

	refracted := reflection.Ray.Direction().InDelta(refractedDirection, core.Tolerance)
	reflected := reflection.Ray.Direction().InDelta(reflectedDirection, core.Tolerance)

	assert.True(t, refracted || reflected)
}

func TestTransparent_ShouldFullyReflect_WhenIncidenceAngleTooLarge(t *testing.T) {
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewTransparent(1.5, randomizer)
	incidentDirection := core.NewVec3(1.0, 1.0, 0.0)
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(1.0, 2.0, 3.0),
		Normal: core.NewVec3(0.0, 1.0, 0.0),
	}

	reflection := material.Reflect(incidentDirection, hit.Point, hit.Normal)

	assert.NotNil(t, reflection)
	assert.Equal(t, color.White, reflection.Attenuation)
	assert.Equal(t, core.NewVec3(1.0, -1.0, 0.0), reflection.Ray.Direction())
	assert.Equal(t, hit.Point, reflection.Ray.Origin())
}
