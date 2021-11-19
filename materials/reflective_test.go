package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestReflective(t *testing.T) {
	material := Reflective{core.Red}
	ray := core.Ray{
		Origin:    core.Vec3{-3.0, 5.0, 3.0},
		Direction: core.Vec3{4.0, -3.0, 0.0},
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.Vec3{1.0, 2.0, 3.0},
		Normal: core.Vec3{0.0, 1.0, 0.0},
	}
	t.Logf("Given a reflective material with color %v and zero fuzziness\n", material.Color)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("\twe can reflect the ray off the material:")
	reflection := material.Reflect(ray, hit)

	expected := Reflection{
		Ray:         core.Ray{hit.Point, core.Vec3{4.0, 3.0, 0.0}},
		Attenuation: material.Color}

	utils.CheckNotNil(t, "reflection", reflection)
	utils.CheckResult(t, "reflection", *reflection, expected)
}

func TestReflective_Reflected(t *testing.T) {
	material := Reflective{core.Red}
	ray := core.Ray{
		Origin:    core.Vec3{-3.0, 5.0, 3.0},
		Direction: core.Vec3{4.0, -3.0, 0.0},
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.Vec3{1.0, 2.0, 3.0},
		Normal: core.Vec3{0.0, 1.0, 0.0},
	}
	t.Logf("Given a reflective material with color %v and zero fuzziness\n", material.Color)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("\twe can reflect the ray off the material:")
	reflection := material.Reflect(ray, hit)

	expected := Reflection{
		Ray:         core.Ray{hit.Point, core.Vec3{4.0, 3.0, 0.0}},
		Attenuation: material.Color}

	utils.CheckNotNil(t, "reflection", reflection)
	utils.CheckResult(t, "reflection", *reflection, expected)
}

func TestReflective_NotReflected(t *testing.T) {
	material := Reflective{core.Red}
	ray := core.Ray{
		Origin:    core.Vec3{-3.0, -1.0, 3.0},
		Direction: core.Vec3{4.0, 3.0, 0.0},
	}
	hit := objects.HitRecord{
		Param:  1.0,
		Point:  core.Vec3{1.0, 2.0, 3.0},
		Normal: core.Vec3{0.0, 1.0, 0.0},
	}
	t.Logf("Given a reflective material with color %v and zero fuzziness\n", material.Color)
	t.Logf("a ray %v and a hit record %v,\n", ray, hit)

	t.Log("\twe should get no reflection as the ray comes from under the surface:")
	reflection := material.Reflect(ray, hit)

	utils.CheckNil(t, "reflection", reflection)
}
