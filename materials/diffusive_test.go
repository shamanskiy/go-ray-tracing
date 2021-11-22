package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestDiffusive_NoRandom(t *testing.T) {
	material := Diffusive{core.Red}
	ray := core.Ray{
		Origin:    core.Vec3{1.0, 2.0, 3.0},
		Direction: core.Vec3{4.0, 5.0, 6.0},
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.Vec3{0.0, 1.0, 2.0},
		Normal: core.Vec3{0.0, 0.0, 1.0},
	}
	t.Logf("Given a diffusive material %v, a ray %v, a hit record %v,\n", material, ray, hit)
	t.Log("and a DISABLED randomizer,")
	core.Random().Disable()
	defer core.Random().Enable()

	t.Log("  we can reflect the ray off the material and expect a predictable result:")
	reflection := material.Reflect(ray, hit)

	expected := Reflection{
		Ray:         core.Ray{hit.Point, hit.Normal},
		Attenuation: material.Color}

	utils.CheckNotNil(t, "reflection", reflection)
	utils.CheckResult(t, "reflection", *reflection, expected)
}

func TestDiffusive_Random(t *testing.T) {
	material := Diffusive{core.Red}
	ray := core.Ray{
		Origin:    core.Vec3{1.0, 2.0, 3.0},
		Direction: core.Vec3{4.0, 5.0, 6.0},
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.Vec3{0.0, 1.0, 2.0},
		Normal: core.Vec3{0.0, 0.0, 1.0},
	}
	t.Logf("Given a diffusive material %v, a ray %v, a hit record %v,\n", material, ray, hit)
	t.Log("and an ENABLED randomizer,")

	t.Log("  the direction of the reflected the ray is random")
	reflection := material.Reflect(ray, hit)
	t.Logf("  but it should be within a unit sphere of the surface normal %v:\n", hit.Normal)

	utils.CheckNotNil(t, "reflection", reflection)
	randomPerturbation := reflection.Ray.Direction.Sub(hit.Normal).Len()
	if randomPerturbation < 1.0 {
		t.Logf("\tPASSED: reflection direction %v, perturbation %v",
			reflection.Ray.Direction, randomPerturbation)
	} else {
		t.Fatalf("\tPASSED: reflection direction %v, perturbation %v",
			reflection.Ray.Direction, randomPerturbation)
	}
}
