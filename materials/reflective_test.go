package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestReflective_Reflected(t *testing.T) {
	material := Reflective{Color: core.Red}
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

	t.Log("  we can reflect the ray off the material and expect a predictable result:")
	reflection := material.Reflect(ray, hit)

	expected := Reflection{
		Ray:         core.Ray{hit.Point, core.Vec3{0.8, 0.6, 0.0}},
		Attenuation: material.Color}

	utils.CheckNotNil(t, "reflection", reflection)
	utils.CheckResult(t, "reflection", *reflection, expected)
}

func TestReflective_NotReflected(t *testing.T) {
	material := Reflective{Color: core.Red}
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

	t.Log("  we should get no reflection as the ray comes from under the surface:")
	reflection := material.Reflect(ray, hit)

	utils.CheckNil(t, "reflection", reflection)
}

func TestReflective_FuzzinessLimits(t *testing.T) {
	t.Log("When we construct a reflective material with fuzziness,")

	t.Log("  if we pass a value between 0 and 1, e.g. 0.5, fuzziness is set to 0.5:")
	material := NewReflectiveWithFuzziness(core.Red, 0.5)
	utils.CheckResult(t, "fuzziness", material.fuzziness, core.Real(0.5))

	t.Log("  if we pass a value less than 0, e.g. -0.2, fuzziness is set to 0.0:")
	material = NewReflectiveWithFuzziness(core.Red, -0.2)
	utils.CheckResult(t, "fuzziness", material.fuzziness, core.Real(0.0))

	t.Log("  if we pass a value greater than 1, e.g. 1.3, fuzziness is set to 1.0:")
	material = NewReflectiveWithFuzziness(core.Red, 1.3)
	utils.CheckResult(t, "fuzziness", material.fuzziness, core.Real(1.0))
}

func TestRefective_WithFuzziness(t *testing.T) {

}
