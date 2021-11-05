package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

func TestDiffusive(t *testing.T) {
	redDiffusive := Diffusive{core.Red}
	ray := core.Ray{
		Origin:    core.Vec3{1.0, 2.0, 3.0},
		Direction: core.Vec3{4.0, 5.0, 6.0},
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.Vec3{0.0, 1.0, 2.0},
		Normal: core.Vec3{0.0, 0.0, 1.0},
	}
	t.Logf("Given a diffusive material %v, a ray %v and a hit record %v,\n", redDiffusive, ray, hit)

	t.Log("we can reflect the ray off the material with disabled randomization:")
	core.Random().Disable()
	defer core.Random().Enable()
	reflection := redDiffusive.Reflect(ray, hit)

	expected := Reflection{
		Ray:         core.Ray{hit.Point, hit.Normal},
		Attenuation: redDiffusive.Color}

	if reflection != nil && *reflection == expected {
		t.Logf("\tPASSED: reflection is %v, expected %v.\n", reflection, expected)
	} else {
		t.Fatalf("\tFAILED: reflection is %v, expected %v.\n", reflection, expected)
	}

}
