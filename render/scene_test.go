package render

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/materials"
	"github.com/Shamanskiy/go-ray-tracer/objects"
	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestScene_Default(t *testing.T) {
	t.Log("Given a default empty scene,")
	scene := Scene{}
	expected := core.Black

	ray := core.Ray{core.Vec3{}, core.Vec3{}}
	t.Logf("\tray %v tests black:\n", ray)
	color := scene.TestRay(ray)
	utils.CheckResult(t, "color", color, expected)

	ray = core.Ray{core.Vec3{1, 2, 3}, core.Vec3{4, 5, 6}}
	t.Logf("\tand any other ray, too, for example %v:\n", ray)
	color = scene.TestRay(ray)
	utils.CheckResult(t, "color", color, expected)
}

func TestScene_Empty(t *testing.T) {
	skyBottom, skyTop := core.Red, core.Blue
	scene := Scene{SkyColorBottom: skyBottom, SkyColorTop: skyTop}
	t.Logf("Given an empty scene with bottom sky color %v and top sky color %v,\n",
		skyBottom, skyTop)

	ray := core.Ray{core.Vec3{0, 0, 0}, core.Vec3{0, -1, 0}}
	t.Logf("\tdown-facing ray %v should return the bottom color:\n", ray)
	color := scene.TestRay(ray)
	expected := skyBottom
	utils.CheckResult(t, "color", color, expected)

	ray = core.Ray{core.Vec3{0, 0, 0}, core.Vec3{0, 1, 0}}
	t.Logf("\tup-facing ray %v should return the top color:\n", ray)
	color = scene.TestRay(ray)
	expected = skyTop
	utils.CheckResult(t, "color", color, expected)

	ray = core.Ray{core.Vec3{0, 0, 0}, core.Vec3{1, 0, 0}}
	t.Logf("\thorizontal ray %v should return the blend color:\n", ray)
	color = scene.TestRay(ray)
	expected = skyTop.Add(skyBottom).Mul(0.5)
	utils.CheckResult(t, "color", color, expected)
}

func TestScene_HitClosetObject(t *testing.T) {
	leftSphere := objects.Sphere{core.Vec3{-6.0, 0.0, 0.0}, 2.0}
	rightSphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 2.0}
	material := materials.Diffusive{core.Black}
	t.Logf("Given a scene with two spheres %v and %v,\n", leftSphere, rightSphere)
	scene := Scene{}
	scene.Add(leftSphere, material)
	scene.Add(rightSphere, material)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("\ta ray with origin %v and direction %v should hit the right sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord, objectIndex := scene.hitClosestObject(hitRay, 0.001)

	expectedHit := objects.HitRecord{Param: 2.0, Point: core.Vec3{2.0, 0.0, 0.0}, Normal: core.Vec3{1.0, 0.0, 0.0}}
	utils.CheckResult(t, "hit record", *hitRecord, expectedHit)
	utils.CheckResult(t, "object index", objectIndex, 1)

	t.Logf("\ta ray with origin %v, direction %v and minimum parameter 7.0 should hit the left sphere:\n", hitRay.Origin, hitRay.Direction)
	hitRecord, objectIndex = scene.hitClosestObject(hitRay, 7.0)

	expectedHit = objects.HitRecord{Param: 8.0, Point: core.Vec3{-4.0, 0.0, 0.0}, Normal: core.Vec3{1.0, 0.0, 0.0}}
	utils.CheckResult(t, "hit record", *hitRecord, expectedHit)
	utils.CheckResult(t, "object index", objectIndex, 0)
}

func TestScene_TestRay_SingleSphere(t *testing.T) {
	t.Log("Given a scene with a blue sphere and a white skybox,")
	skyColor := core.White
	scene := Scene{SkyColorBottom: skyColor, SkyColorTop: skyColor}
	sphere := objects.Sphere{core.Vec3{0.0, 0.0, 0.0}, 1.0}
	sphereColor := core.Blue
	material := materials.Diffusive{sphereColor}
	scene.Add(sphere, material)

	hitRay := core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{-1.0, 0.0, 0.0}}
	t.Logf("\ta ray with origin %v and direction %v should return blue color:\n", hitRay.Origin, hitRay.Direction)
	rayColor := scene.TestRay(hitRay)
	expectedColor := sphereColor
	utils.CheckResult(t, "ray color", rayColor, expectedColor)

	hitRay = core.Ray{core.Vec3{4.0, 0.0, 0.0}, core.Vec3{0.0, 1.0, 0.0}}
	t.Logf("\ta ray with origin %v and direction %v should return white color:\n", hitRay.Origin, hitRay.Direction)
	rayColor = scene.TestRay(hitRay)
	expectedColor = skyColor
	utils.CheckResult(t, "ray color", rayColor, expectedColor)
}
