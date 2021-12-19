package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestTransparent_RefactionIndexLimits(t *testing.T) {
	t.Log("When we construct a transparent material,")

	t.Log("  if we pass a refractive index less than 1, e.g. 0, the index is set to 1:")
	material := NewTransparent(0.0)
	utils.CheckResult(t, "refractive index", material.refractionIndex, core.Real(1))

	t.Log("  if we pass a refractive index equal or greater than 0.001, e.g. 1.5, the index is set to 1.5:")
	material = NewTransparent(1.5)
	utils.CheckResult(t, "refractive index", material.refractionIndex, core.Real(1.5))
}

func TestRefract_OutsideToInside_Glass(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.5)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{1., -1., 0.}
	t.Logf("  incident vector %v gets refracted:\n", incidentVector)

	refractedDirection, cosOut := refract(incidentVector.Normalize(), normal, 1./refractionIndex)
	utils.CheckResult(t, "Refracted angle cosine", cosOut, core.Real(0.8819171))
	utils.CheckNotNil(t, "Refracted direction", refractedDirection)
	utils.CheckResult(t, "Refracted direction", *refractedDirection, core.Vec3{0.47140452, -0.8819171, 0})
}

func TestRefract_OutsideToInside_Air(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{1., -1., 0.}
	t.Logf("  incident vector %v doesn't get refracted:\n", incidentVector)

	refractedDirection, cosOut := refract(incidentVector.Normalize(), normal, 1./refractionIndex)
	utils.CheckResult(t, "Refracted angle cosine", cosOut, core.Real(0.70710677))
	utils.CheckNotNil(t, "Refracted direction", refractedDirection)
	utils.CheckResult(t, "Refracted direction", *refractedDirection, core.Vec3{0.70710677, -0.70710677, 0})
}

func TestRefract_InsideToOutside_Glass(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.5)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{0.6, 0.8, 0.}
	t.Logf("  incident vector %v gets refracted:\n", incidentVector)

	refractedDirection, cosOut := refract(incidentVector.Normalize(), normal, refractionIndex)
	core.CheckFloatSameTol(t, "Refracted angle cosine", cosOut, core.Real(0.435889))
	utils.CheckNotNil(t, "Refracted direction", refractedDirection)
	core.CheckVec3SameTol(t, "Refracted direction", *refractedDirection, core.Vec3{0.9, 0.435889, 0})
}

func TestRefract_InsideToOutside_Air(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{0.6, 0.8, 0.}
	t.Logf("  incident vector %v doesn't get refracted:\n", incidentVector)

	refractedDirection, cosOut := refract(incidentVector.Normalize(), normal, refractionIndex)
	utils.CheckResult(t, "Refracted angle cosine", cosOut, core.Real(0.8))
	utils.CheckNotNil(t, "Refracted direction", refractedDirection)
	utils.CheckResult(t, "Refracted direction", *refractedDirection, core.Vec3{0.6, 0.8, 0})
}

func TestRefract_InsideToOutside_Glass_FullReflection(t *testing.T) {
	normal := core.Vec3{0., 1., 0.}
	refractionIndex := core.Real(1.5)
	t.Logf("Given a material with refraction index %v and normal %v,\n", refractionIndex, normal)

	incidentVector := core.Vec3{1., 1., 0.}
	t.Logf("  incident vector %v gets fully reflected because the incident angle is too big:\n", incidentVector)

	refractedDirection, cosOut := refract(incidentVector.Normalize(), normal, refractionIndex)
	utils.CheckResult(t, "Refracted angle cosine", cosOut, core.Real(0.))
	utils.CheckNil(t, "Refracted direction", refractedDirection)
}

func TestSchlickLaw(t *testing.T) {
	refractionIndex := core.Real(1.5)
	t.Logf("Given a material with refraction index %v,\n", refractionIndex)

	cosIn := core.Real(1.0)
	t.Logf("  for an incident angle cosine  %v, we can compute the reflection ratio:\n", cosIn)

	reflectionRatio := schlickLaw(cosIn, refractionIndex)
	core.CheckFloatSameTol(t, "Reflected ratio", reflectionRatio, core.Real(0.04))

	cosIn = core.Real(0.0)
	t.Logf("  for an incident angle cosine  %v, we can compute the reflection ratio:\n", cosIn)

	reflectionRatio = schlickLaw(cosIn, refractionIndex)
	core.CheckFloatSameTol(t, "Reflected ratio", reflectionRatio, core.Real(1))

	cosIn = core.Real(0.5)
	t.Logf("  for an incident angle cosine  %v, we can compute the reflection ratio:\n", cosIn)

	reflectionRatio = schlickLaw(cosIn, refractionIndex)
	core.CheckFloatSameTol(t, "Reflected ratio", reflectionRatio, core.Real(0.07))
}
