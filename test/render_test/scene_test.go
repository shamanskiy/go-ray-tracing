package render

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/Shamanskiy/go-ray-tracer/src/render"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/chewxy/math32"
)

func TestScene_Default(t *testing.T) {
	t.Log("Given a default empty scene,")
	scene := render.Scene{}
	expected := core.Black

	ray := core.Ray{core.Vec3{}, core.Vec3{}}
	t.Logf("  ray %v tests black:\n", ray)
	color := scene.TestRay(ray)
	test.CheckResult(t, "color", color, expected)

	ray = core.Ray{core.Vec3{1, 2, 3}, core.Vec3{4, 5, 6}}
	t.Logf("  and any other ray, too, for example %v:\n", ray)
	color = scene.TestRay(ray)
	test.CheckResult(t, "color", color, expected)
}

func TestScene_Empty(t *testing.T) {
	skyBottom, skyTop := core.Red, core.Blue
	scene := render.Scene{SkyColorBottom: skyBottom, SkyColorTop: skyTop}
	t.Logf("Given an empty scene with bottom sky color %v and top sky color %v,\n",
		skyBottom, skyTop)

	ray := core.Ray{core.Vec3{0, 0, 0}, core.Vec3{0, -1, 0}}
	t.Logf("  down-facing ray %v should return the bottom color:\n", ray)
	color := scene.TestRay(ray)
	expected := skyBottom
	test.CheckResult(t, "color", color, expected)

	ray = core.Ray{core.Vec3{0, 0, 0}, core.Vec3{0, 1, 0}}
	t.Logf("  up-facing ray %v should return the top color:\n", ray)
	color = scene.TestRay(ray)
	expected = skyTop
	test.CheckResult(t, "color", color, expected)

	ray = core.Ray{core.Vec3{0, 0, 0}, core.Vec3{1, 0, 0}}
	t.Logf("  horizontal ray %v should return the blend color:\n", ray)
	color = scene.TestRay(ray)
	expected = skyTop.Add(skyBottom).Mul(0.5)
	test.CheckResult(t, "color", color, expected)
}

func TestScene_HitClosetObject(t *testing.T) {
	leftSphere := objects.Sphere{core.Vec3{-6.0, 0.0, 0.0}, 2.0}
	rightSphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	material := materials.Diffusive{core.Black}
	t.Logf("Given a scene with two spheres %v and %v,\n", leftSphere, rightSphere)
	scene := render.Scene{}
	scene.Add(leftSphere, material)
	scene.Add(rightSphere, material)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("  a ray with origin %v and direction %v should hit the right sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord, objectIndex := scene.HitClosestObject(hitRay, 0.001)

	expectedHit := objects.HitRecord{Param: 2.0, Point: core.Vec3{2.0, 0.0, 0.0}, Normal: core.Vec3{1.0, 0.0, 0.0}}
	test.CheckResult(t, "hit record", *hitRecord, expectedHit)
	test.CheckResult(t, "object index", objectIndex, 1)

	t.Logf("  a ray with origin %v, direction %v and minimum parameter 7.0 should hit the left sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord, objectIndex = scene.HitClosestObject(hitRay, 7.0)

	expectedHit = objects.HitRecord{Param: 8.0, Point: core.Vec3{-4.0, 0.0, 0.0}, Normal: core.Vec3{1.0, 0.0, 0.0}}
	test.CheckResult(t, "hit record", *hitRecord, expectedHit)
	test.CheckResult(t, "object index", objectIndex, 0)
}

func TestScene_TestRay_SingleSphere(t *testing.T) {
	t.Log("Given a scene with a blue sphere and a white skybox,")
	skyColor := core.White
	scene := render.Scene{SkyColorBottom: skyColor, SkyColorTop: skyColor}
	sphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 1.0}
	sphereColor := core.Blue
	material := materials.Diffusive{sphereColor}
	scene.Add(sphere, material)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("  a ray with origin %v and direction %v should return blue color:\n", hitRay.Origin, hitRay.Direction)
	rayColor := scene.TestRay(hitRay)
	expectedColor := sphereColor
	test.CheckResult(t, "ray color", rayColor, expectedColor)

	hitRay = core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{0.0, 1.0, 0.0}}
	t.Logf("  a ray with origin %v and direction %v should return white color:\n", hitRay.Origin, hitRay.Direction)
	rayColor = scene.TestRay(hitRay)
	expectedColor = skyColor
	test.CheckResult(t, "ray color", rayColor, expectedColor)
}

func TestScene_NumberOfReflectionsExceeded(t *testing.T) {
	t.Log("Given a scene with a white skybox,")
	skyColor := core.White
	scene := render.Scene{SkyColorBottom: skyColor, SkyColorTop: skyColor}
	t.Log("two red spheres,")
	sphereA := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 1.0}
	sphereB := objects.Sphere{core.Vec3{4.0, 0.0, 0.0}, 1.0}
	material := materials.Diffusive{core.Red}
	scene.Add(sphereA, material)
	scene.Add(sphereB, material)

	t.Log("and a disabled randomizer,")
	core.Random().Disable()
	defer core.Random().Enable()

	ray := core.Ray{core.Vec3{2.0, 0.0, 0.0}, core.Vec3{1.0, 0.0, 0.0}}
	t.Log("  a ray colinear with the line between the spheres' centers")
	t.Log("  will bounce between the spheres until the number of reflections is exceeded.")
	t.Log("  The resulting color should be black:")
	rayColor := scene.TestRay(ray)
	expected := core.Black

	test.CheckResult(t, "ray color", rayColor, expected)
}

func TestScene_TwoReflections(t *testing.T) {
	t.Log("Given a scene with a white skybox,")
	skyColor := core.White
	scene := render.Scene{SkyColorBottom: skyColor, SkyColorTop: skyColor}
	t.Log("a disabled randomizer,")
	core.Random().Disable()
	defer core.Random().Enable()

	sphereA := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 1.0}
	sphereB := objects.Sphere{core.Vec3{4.0, 3.0, 0.0}, 1.0}
	sphereColor := core.Vec3{0.2, 0.4, 0.6}
	material := materials.Diffusive{sphereColor}
	scene.Add(sphereA, material)
	scene.Add(sphereB, material)
	t.Logf("and two spheres %v and %v with diffusive color %v,\n",
		sphereA, sphereB, sphereColor)

	firstHitPoint := core.Vec3{math32.Sqrt(2) / 2, math32.Sqrt(2) / 2, 0.0}
	rayOrigin := core.Vec3{10.0, 0.0, 0.0}
	ray := core.Ray{rayOrigin, firstHitPoint.Sub(rayOrigin)}
	t.Logf("  a ray %v should hit the first sphere at point %v,\n", ray, firstHitPoint)
	t.Log("  get reflected and hit the second sphere at point [3.0 3.0 0.0],")
	t.Log("  then get reflected towards the sky:")
	rayColor := scene.TestRay(ray)
	expected := core.MulElem(sphereColor, sphereColor)

	test.CheckResult(t, "ray color", rayColor, expected)
}
