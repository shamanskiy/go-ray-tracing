package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/stretchr/testify/assert"
)

func TestDiffusive_NoRandom(t *testing.T) {
	material := materials.Diffusive{color.Red}
	ray := core.NewRay(core.NewVec3(1.0, 2.0, 3.0), core.NewVec3(4.0, 5.0, 6.0))
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(0.0, 1.0, 2.0),
		Normal: core.NewVec3(0.0, 0.0, 1.0),
	}
	t.Logf("Given a diffusive material %v, a ray %v, a hit record %v,\n", material, ray, hit)
	t.Log("and a DISABLED randomizer,")
	core.Random().Disable()
	defer core.Random().Enable()

	t.Log("  we can reflect the ray off the material and expect a predictable result:")
	reflection := material.Reflect(ray, hit)

	expected := materials.Reflection{
		Ray:         core.NewRay(hit.Point, hit.Normal),
		Attenuation: material.Color}

	assert.Equal(t, expected, *reflection)
}

func TestDiffusive_Random(t *testing.T) {
	material := materials.Diffusive{color.Red}
	ray := core.NewRay(core.NewVec3(1.0, 2.0, 3.0), core.NewVec3(4.0, 5.0, 6.0))
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.NewVec3(0.0, 1.0, 2.0),
		Normal: core.NewVec3(0.0, 0.0, 1.0),
	}
	t.Logf("Given a diffusive material %v, a ray %v, a hit record %v,\n", material, ray, hit)
	t.Log("and an ENABLED randomizer,")

	t.Log("  the direction of the reflected the ray is random")
	reflection := material.Reflect(ray, hit)
	t.Logf("  but it should be within a unit sphere of the surface normal %v:\n", hit.Normal)

	assert.NotNil(t, reflection)
	randomPerturbation := reflection.Ray.Direction().Sub(hit.Normal).Len()
	if randomPerturbation < 1.0 {
		t.Logf("\tPASSED: reflection direction %v, perturbation %v",
			reflection.Ray.Direction, randomPerturbation)
	} else {
		t.Fatalf("\tPASSED: reflection direction %v, perturbation %v",
			reflection.Ray.Direction, randomPerturbation)
	}
}
