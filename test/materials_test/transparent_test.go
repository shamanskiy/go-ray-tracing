package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/Shamanskiy/go-ray-tracer/test"
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
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{1., -1., 0.}
	t.Logf("  incident vector %v gets refracted:\n", incidentVector)
	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.CheckNotNil(t, "Refracted direction", refractedDirection)
	core.CheckVec3Tol(t, "Refracted direction", *refractedDirection, core.Vec3{0.666666, -1.247219, 0})
	core.CheckRealTol(t, "Reflection ratio", reflectionRatio, core.Real(0.04207))

	incidentVector = core.Vec3{1., -0.1, 0.}
	t.Logf("  incident vector %v gets refracted:\n", incidentVector)
	refractedDirection, reflectionRatio = materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.CheckNotNil(t, "Refracted direction", refractedDirection)
	core.CheckVec3Tol(t, "Refracted direction", *refractedDirection, core.Vec3{0.666666, -0.752034, 0})
	core.CheckRealTol(t, "Reflection ratio", reflectionRatio, core.Real(0.60843))

	incidentVector = core.Vec3{1., -0.01, 0.}
	t.Logf("  incident vector %v gets refracted:\n", incidentVector)
	refractedDirection, reflectionRatio = materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.CheckNotNil(t, "Refracted direction", refractedDirection)
	core.CheckVec3Tol(t, "Refracted direction", *refractedDirection, core.Vec3{0.666666, -0.745423, 0})
	core.CheckRealTol(t, "Reflection ratio", reflectionRatio, core.Real(0.95295))
}

func TestRefract_OutsideToInside_Air(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{1., -1., 0.}
	t.Logf("  incident vector %v doesn't get refracted:\n", incidentVector)

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.CheckNotNil(t, "Refracted direction", refractedDirection)
	core.CheckVec3Tol(t, "Refracted direction", *refractedDirection, core.Vec3{1., -1., 0})
	core.CheckRealTol(t, "Reflection ratio", reflectionRatio, core.Real(0.002155))
}

func TestRefract_InsideToOutside_Glass(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.5)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{3, 4, 0.}
	t.Logf("  incident vector %v gets refracted:\n", incidentVector)

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.CheckNotNil(t, "Refracted direction", refractedDirection)
	core.CheckVec3Tol(t, "Refracted direction", *refractedDirection, core.Vec3{4.5, 2.179448, 0})
	core.CheckRealTol(t, "Reflection ratio", reflectionRatio, core.Real(0.0948391))
}

func TestRefract_InsideToOutside_Air(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{3, 4, 0.}
	t.Logf("  incident vector %v doesn't get refracted:\n", incidentVector)

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.CheckNotNil(t, "Refracted direction", refractedDirection)
	core.CheckVec3Tol(t, "Refracted direction", *refractedDirection, core.Vec3{3, 4, 0})
	core.CheckRealTol(t, "Reflection ratio", reflectionRatio, core.Real(0.00032))
}

func TestRefract_InsideToOutside_Glass_FullReflection(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.5)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{1., 1., 0.}
	t.Logf("  incident vector %v gets fully reflected because the incident angle is too big:\n", incidentVector)

	refractedDirection, reflectionRatio := materials.ComputeRefraction(incidentVector, normal, refractionIndex)
	test.CheckNil(t, "Refracted direction", refractedDirection)
	core.CheckRealTol(t, "Reflection ratio", reflectionRatio, core.Real(1.0))
}

func TestSchlickLaw(t *testing.T) {
	refractionIndex := core.Real(1.5)
	t.Logf("Given a material with refraction index %v,\n", refractionIndex)

	cosIn := core.Real(1.0)
	t.Logf("  for an incident angle cosine  %v, we can compute the reflection ratio:\n", cosIn)
	reflectionRatio := materials.SchlickLaw(cosIn, refractionIndex)
	core.CheckRealTol(t, "Reflected ratio", reflectionRatio, core.Real(0.04))

	cosIn = core.Real(0.0)
	t.Logf("  for an incident angle cosine  %v, we can compute the reflection ratio:\n", cosIn)
	reflectionRatio = materials.SchlickLaw(cosIn, refractionIndex)
	core.CheckRealTol(t, "Reflected ratio", reflectionRatio, core.Real(1))

	cosIn = core.Real(0.5)
	t.Logf("  for an incident angle cosine  %v, we can compute the reflection ratio:\n", cosIn)
	reflectionRatio = materials.SchlickLaw(cosIn, refractionIndex)
	core.CheckRealTol(t, "Reflected ratio", reflectionRatio, core.Real(0.07))
}

func TestTransparent(t *testing.T) {
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
	t.Logf("Given a transparent material with refraction index %v,\n", material.RefractionIndex)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("  the ray either gets refracted or reflected, and the attenuation is white:")
	reflection := material.Reflect(ray, hit)
	test.CheckNotNil(t, "reflection", reflection)
	test.CheckResult(t, "Refraction color", reflection.Attenuation, core.White)
	test.CheckResult(t, "Ray origin", reflection.Ray.Origin, hit.Point)

	refractedDirection := core.Vec3{0.666666, -1.247219, 0}
	reflectedDirection := core.Vec3{1.0, 1.0, 0.0}

	t.Logf("  expected refraction direction: %v", refractedDirection)
	t.Logf("  expected reflection direction: %v", reflectedDirection)

	if core.IsSameVec3(reflection.Ray.Direction, refractedDirection) ||
		core.IsSameVec3(reflection.Ray.Direction, reflectedDirection) {
		t.Logf("\tPASSED: Ray direction %v", reflection.Ray.Direction)
	} else {
		t.Fatalf("\tPASSED: Ray direction %v", reflection.Ray.Direction)
	}
}

func TestTransparent_FullReflection(t *testing.T) {
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
	t.Logf("Given a transparent material with refraction index %v,\n", material.RefractionIndex)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("  the ray gets fully reflected because the incidence angle is too large:")
	reflection := material.Reflect(ray, hit)
	test.CheckNotNil(t, "reflection", reflection)
	test.CheckResult(t, "Refraction color", reflection.Attenuation, core.White)
	test.CheckResult(t, "Ray origin", reflection.Ray.Origin, hit.Point)
	expectedDirection := core.Vec3{1.0, -1.0, 0.0}
	test.CheckResult(t, "Ray direction", reflection.Ray.Direction, expectedDirection)
}
