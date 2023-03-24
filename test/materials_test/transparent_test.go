package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

func TestTransparent_RefactionIndexLimits(t *testing.T) {
	t.Log("When we construct a transparent material,")

	t.Log("  if we pass a refractive index less than 1, e.g. 0, the index is set to 1:")
	material := materials.NewTransparent(0.0)
	test.CheckResult(t, "refractive index", material.RefractionIndex, core.Real(1))

	t.Log("  if we pass a refractive index equal or greater than 0.001, e.g. 1.5, the index is set to 1.5:")
	material = materials.NewTransparent(1.5)
	test.CheckResult(t, "refractive index", material.RefractionIndex, core.Real(1.5))
}

func TestRefract_OutsideToInside_Glass(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.5)

	incidentVector := core.Vec3{1., -1., 0.}
	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.Vec3{0.666666, -1.247219, 0}, *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.04207, reflectionRatio, core.Tolerance)

	incidentVector = core.Vec3{1., -0.1, 0.}
	refractedDirection, reflectionRatio = materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.Vec3{0.666666, -0.752034, 0}, *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.60843, reflectionRatio, core.Tolerance)

	incidentVector = core.Vec3{1., -0.01, 0.}
	refractedDirection, reflectionRatio = materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.Vec3{0.666666, -0.745423, 0}, *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.95295, reflectionRatio, core.Tolerance)
}

func TestRefract_OutsideToInside_Air(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.)

	incidentVector := core.Vec3{1., -1., 0.}

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.Vec3{1, -1, 0}, *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.002155, reflectionRatio, core.Tolerance)
}

func TestRefract_InsideToOutside_Glass(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.5)

	incidentVector := core.Vec3{3, 4, 0.}

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)

	test.AssertInDeltaVec3(t, core.Vec3{4.5, 2.179448, 0}, *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.0948391, reflectionRatio, core.Tolerance)
}

func TestRefract_InsideToOutside_Air(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.)

	incidentVector := core.Vec3{3, 4, 0.}

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.AssertInDeltaVec3(t, core.Vec3{3, 4, 0}, *refractedDirection, core.Tolerance)
	assert.InDelta(t, 0.00032, reflectionRatio, core.Tolerance)
}

func TestRefract_InsideToOutside_Glass_ShouldFullyReflect_AngleTooBig(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.5)

	incidentVector := core.Vec3{1., 1., 0.}

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
	material := materials.NewTransparent(1.5)
	ray := core.Ray{
		Origin:    core.Vec3{0.0, 0.0, 3.0},
		Direction: core.Vec3{1.0, -1.0, 0.0},
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.Vec3{1.0, 2.0, 3.0},
		Normal: core.Vec3{0.0, 1.0, 0.0},
	}

	reflection := material.Reflect(ray, hit)

	assert.NotNil(t, reflection)
	assert.Equal(t, core.White, reflection.Attenuation)
	assert.Equal(t, hit.Point, reflection.Ray.Origin)

	refractedDirection := core.Vec3{0.666666, -1.247219, 0}
	reflectedDirection := core.Vec3{1.0, 1.0, 0.0}

	refracted := core.IsVec3InDelta(reflection.Ray.Direction, refractedDirection, core.Tolerance)
	reflected := core.IsVec3InDelta(reflection.Ray.Direction, reflectedDirection, core.Tolerance)

	assert.True(t, refracted || reflected)
}

func TestTransparent_ShouldFullyReflect_WhenIncidenceAngleTooLarge(t *testing.T) {
	material := materials.NewTransparent(1.5)
	ray := core.Ray{
		Origin:    core.Vec3{0.0, 0.0, 3.0},
		Direction: core.Vec3{1.0, 1.0, 0.0},
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.Vec3{1.0, 2.0, 3.0},
		Normal: core.Vec3{0.0, 1.0, 0.0},
	}

	reflection := material.Reflect(ray, hit)

	assert.NotNil(t, reflection)
	assert.Equal(t, core.White, reflection.Attenuation)
	assert.Equal(t, core.Vec3{1.0, -1.0, 0.0}, reflection.Ray.Direction)
	assert.Equal(t, hit.Point, reflection.Ray.Origin)
}
