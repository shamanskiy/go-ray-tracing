package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/stretchr/testify/assert"
)

func TestReflective_Reflected(t *testing.T) {
	material := materials.Reflective{Color: color.Red}
	ray := core.Ray{
		Origin:    core.NewVec3(-3.0, 5.0, 3.0),
		Direction: core.NewVec3(4.0, -3.0, 0.0),
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(1.0, 2.0, 3.0),
		Normal: core.NewVec3(0.0, 1.0, 0.0),
	}
	t.Logf("Given a reflective material with color %v and zero fuzziness\n", material.Color)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("  we can reflect the ray off the material and expect a predictable result:")
	reflection := material.Reflect(ray, hit)

	expected := materials.Reflection{
		Ray:         core.Ray{hit.Point, core.NewVec3(0.8, 0.6, 0.0)},
		Attenuation: material.Color}

	assert.Equal(t, expected, *reflection)
}

func TestReflective_NotReflected(t *testing.T) {
	material := materials.Reflective{Color: color.Red}
	ray := core.Ray{
		Origin:    core.NewVec3(-3.0, -1.0, 3.0),
		Direction: core.NewVec3(4.0, 3.0, 0.0),
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(1.0, 2.0, 3.0),
		Normal: core.NewVec3(0.0, 1.0, 0.0),
	}
	t.Logf("Given a reflective material with color %v and zero fuzziness\n", material.Color)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("  we should get no reflection as the ray comes from under the surface:")
	reflection := material.Reflect(ray, hit)

	assert.Nil(t, reflection)
}

func TestReflective_FuzzinessLimits(t *testing.T) {
	t.Log("When we construct a reflective material with fuzziness,")

	t.Log("  if we pass a value between 0 and 1, e.g. 0.5, fuzziness is set to 0.5:")
	material := materials.NewReflectiveFuzzy(color.Red, 0.5)
	assert.EqualValues(t, 0.5, material.Fuzziness)

	t.Log("  if we pass a value less than 0, e.g. -0.2, fuzziness is set to 0.0:")
	material = materials.NewReflectiveFuzzy(color.Red, -0.2)
	assert.EqualValues(t, 0, material.Fuzziness)

	t.Log("  if we pass a value greater than 1, e.g. 1.3, fuzziness is set to 1.0:")
	material = materials.NewReflectiveFuzzy(color.Red, 1.3)
	assert.EqualValues(t, 1, material.Fuzziness)
}

func TestRefective_WithFuzziness(t *testing.T) {
	material := materials.NewReflectiveFuzzy(color.Red, 0.5)
	ray := core.Ray{
		Origin:    core.NewVec3(-3.0, 5.0, 3.0),
		Direction: core.NewVec3(4.0, -3.0, 0.0),
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(1.0, 2.0, 3.0),
		Normal: core.NewVec3(0.0, 1.0, 0.0),
	}
	t.Logf("Given a reflective material with color %v and fuzziness %v,\n", material.Color, material.Fuzziness)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("  the direction of the reflected the ray is random")
	reflection := material.Reflect(ray, hit)
	expectedMeanDirection := core.NewVec3(0.8, 0.6, 0.0)
	t.Logf("  but it should be within a sphere of radius %v of the expected direction %v:\n",
		material.Fuzziness, expectedMeanDirection)

	assert.NotNil(t, reflection)
	randomPerturbation := reflection.Ray.Direction.Sub(expectedMeanDirection).Len()
	if randomPerturbation < material.Fuzziness {
		t.Logf("\tPASSED: reflection direction %v, perturbation %v",
			reflection.Ray.Direction, randomPerturbation)
	} else {
		t.Fatalf("\tPASSED: reflection direction %v, perturbation %v",
			reflection.Ray.Direction, randomPerturbation)
	}
}
