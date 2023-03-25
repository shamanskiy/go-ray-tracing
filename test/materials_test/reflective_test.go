package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/stretchr/testify/assert"
)

func TestReflective_Reflected(t *testing.T) {
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewReflective(color.Red, randomizer)
	ray := core.NewRay(core.NewVec3(-3.0, 5.0, 3.0), core.NewVec3(4.0, -3.0, 0.0))
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(1.0, 2.0, 3.0),
		Normal: core.NewVec3(0.0, 1.0, 0.0),
	}
	t.Logf("Given a reflective material with color %v and zero fuzziness\n", color.Red)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("  we can reflect the ray off the material and expect a predictable result:")
	reflection := material.Reflect(ray, hit)

	expected := materials.Reflection{
		Ray:         core.NewRay(hit.Point, core.NewVec3(0.8, 0.6, 0.0)),
		Attenuation: color.Red}

	assert.Equal(t, expected, *reflection)
}

func TestReflective_NotReflected(t *testing.T) {
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewReflective(color.Red, randomizer)
	ray := core.NewRay(core.NewVec3(-3.0, -1.0, 3.0), core.NewVec3(4.0, 3.0, 0.0))
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(1.0, 2.0, 3.0),
		Normal: core.NewVec3(0.0, 1.0, 0.0),
	}
	t.Logf("Given a reflective material with color %v and zero fuzziness\n", color.Red)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("  we should get no reflection as the ray comes from under the surface:")
	reflection := material.Reflect(ray, hit)

	assert.Nil(t, reflection)
}

func TestRefective_WithFuzziness(t *testing.T) {
	fuzziness := float32(0.5)
	materialColor := color.Red
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewReflectiveFuzzy(materialColor, 0.5, randomizer)
	ray := core.NewRay(core.NewVec3(-3.0, 5.0, 3.0), core.NewVec3(4.0, -3.0, 0.0))
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(1.0, 2.0, 3.0),
		Normal: core.NewVec3(0.0, 1.0, 0.0),
	}
	t.Logf("Given a reflective material with color %v and fuzziness %v,\n", materialColor, fuzziness)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("  the direction of the reflected the ray is random")
	reflection := material.Reflect(ray, hit)
	expectedMeanDirection := core.NewVec3(0.8, 0.6, 0.0)
	t.Logf("  but it should be within a sphere of radius %v of the expected direction %v:\n",
		fuzziness, expectedMeanDirection)

	assert.NotNil(t, reflection)
	randomPerturbation := reflection.Ray.Direction().Sub(expectedMeanDirection).Len()
	if randomPerturbation < fuzziness {
		t.Logf("\tPASSED: reflection direction %v, perturbation %v",
			reflection.Ray.Direction(), randomPerturbation)
	} else {
		t.Fatalf("\tPASSED: reflection direction %v, perturbation %v",
			reflection.Ray.Direction(), randomPerturbation)
	}
}
