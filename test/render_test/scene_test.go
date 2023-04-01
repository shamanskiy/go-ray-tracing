package render

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/Shamanskiy/go-ray-tracer/src/render"
	"github.com/chewxy/math32"
	"github.com/stretchr/testify/assert"
)

var anyPoint = core.NewVec3(999, 666, 333)
var anyDirection = core.NewVec3(1, 2, 3)
var anyObjectColor = color.Red
var anyOtherObjectColor = color.Green
var anyBackgroundColor = color.Blue

func TestScene_ShouldReturnBackgroundColor_IfSceneEmpty(t *testing.T) {
	flatColor := background.NewFlatColor(anyBackgroundColor)
	scene := render.NewScene(flatColor)
	ray := core.NewRay(anyPoint, anyDirection)

	rayColor := scene.TestRay(ray)

	assert.Equal(t, flatColor.Color(), rayColor)
}

func TestScene_ShouldReturnObjectColorMixedWithBackgroundColor_IfObjectHitOnce(t *testing.T) {
	background := background.NewFlatColor(anyBackgroundColor)
	scene := render.NewScene(background)
	sphere := objects.NewSphere(core.NewVec3(0, 0, 0), 1)
	material := materials.NewDiffusive(anyObjectColor, random.NewRandomGenerator())
	scene.Add(sphere, material)
	ray := core.NewRay(core.NewVec3(2, 0, 0), core.NewVec3(-1, 0, 0))

	rayColor := scene.TestRay(ray)

	assert.Equal(t, material.Color().MulColor(background.Color()), rayColor)
}

func TestScene_ShouldHitClosestObject(t *testing.T) {
	background := background.NewFlatColor(anyBackgroundColor)
	scene := render.NewScene(background)
	sphere1 := objects.NewSphere(core.NewVec3(0, 0, 0), 1)
	material1 := materials.NewDiffusive(anyObjectColor, random.NewRandomGenerator())
	scene.Add(sphere1, material1)
	sphere2 := objects.NewSphere(core.NewVec3(-10, 0, 0), 1)
	material2 := materials.NewDiffusive(anyOtherObjectColor, random.NewRandomGenerator())
	ray := core.NewRay(core.NewVec3(2, 0, 0), core.NewVec3(-1, 0, 0))
	scene.Add(sphere2, material2)

	rayColor := scene.TestRay(ray)

	assert.Equal(t, material1.Color().MulColor(background.Color()), rayColor)
}

func TestScene_HitClosetObject(t *testing.T) {
	leftSphere := objects.NewSphere(core.NewVec3(-6.0, 0.0, 0.0), 2)
	rightSphere := objects.NewSphere(core.NewVec3(0.0, 0.0, 0.0), 2)
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewDiffusive(color.Black, randomizer)
	t.Logf("Given a scene with two spheres %v and %v,\n", leftSphere, rightSphere)
	scene := render.Scene{}
	scene.Add(leftSphere, material)
	scene.Add(rightSphere, material)

	hitRay := core.NewRay(core.NewVec3(4.0, 0.0, 0.0), core.NewVec3(-1.0, 0.0, 0.0))
	t.Logf("  a ray with origin %v and direction %v should hit the right sphere:\n", hitRay.Origin(), hitRay.Direction())
	scene.SetMinRayHitParameter(0.001)
	hitRecord, objectIndex := scene.HitClosestObject(hitRay)

	// expectedHit := objects.HitRecord{Param: 2.0, Point: core.NewVec3(2.0, 0.0, 0.0), Normal: core.NewVec3(1.0, 0.0, 0.0)}
	expectedHit := objects.HitRecord{Point: core.NewVec3(2.0, 0.0, 0.0), Normal: core.NewVec3(1.0, 0.0, 0.0)}
	assert.Equal(t, expectedHit, *hitRecord)
	assert.Equal(t, 1, objectIndex)

	t.Logf("  a ray with origin %v, direction %v and minimum parameter 7.0 should hit the left sphere:\n", hitRay.Origin(), hitRay.Direction())
	scene.SetMinRayHitParameter(7.0)
	hitRecord, objectIndex = scene.HitClosestObject(hitRay)

	expectedHit = objects.HitRecord{Point: core.NewVec3(-4.0, 0.0, 0.0), Normal: core.NewVec3(1.0, 0.0, 0.0)}
	// expectedHit = objects.HitRecord{Param: 8.0, Point: core.NewVec3(-4.0, 0.0, 0.0), Normal: core.NewVec3(1.0, 0.0, 0.0)}
	assert.Equal(t, expectedHit, *hitRecord)
	assert.Equal(t, 0, objectIndex)
}

func TestScene_TestRay_SingleSphere(t *testing.T) {
	t.Log("Given a scene with a blue sphere and a white skybox,")
	skyColor := color.White
	scene := render.NewScene(background.NewFlatColor(skyColor))
	sphere := objects.NewSphere(core.NewVec3(0.0, 0.0, 0.0), 1)
	sphereColor := color.Blue
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewDiffusive(sphereColor, randomizer)
	scene.Add(sphere, material)

	hitRay := core.NewRay(core.NewVec3(4.0, 0.0, 0.0), core.NewVec3(-1.0, 0.0, 0.0))
	t.Logf("  a ray with origin %v and direction %v should return blue color:\n", hitRay.Origin(), hitRay.Direction())
	rayColor := scene.TestRay(hitRay)
	expectedColor := sphereColor
	assert.Equal(t, expectedColor, rayColor)

	hitRay = core.NewRay(core.NewVec3(4.0, 0.0, 0.0), core.NewVec3(0.0, 1.0, 0.0))
	t.Logf("  a ray with origin %v and direction %v should return white color:\n", hitRay.Origin(), hitRay.Direction())
	rayColor = scene.TestRay(hitRay)
	expectedColor = skyColor

	assert.Equal(t, expectedColor, rayColor)
}

func TestScene_NumberOfReflectionsExceeded(t *testing.T) {
	t.Log("Given a scene with a white skybox,")
	skyColor := color.White
	scene := render.NewScene(background.NewFlatColor(skyColor))
	t.Log("two red spheres,")
	sphereA := objects.NewSphere(core.NewVec3(0.0, 0.0, 0.0), 1)
	sphereB := objects.NewSphere(core.NewVec3(4.0, 0.0, 0.0), 1)
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewDiffusive(color.Red, randomizer)
	scene.Add(sphereA, material)
	scene.Add(sphereB, material)

	ray := core.NewRay(core.NewVec3(2.0, 0.0, 0.0), core.NewVec3(1.0, 0.0, 0.0))
	t.Log("  a ray colinear with the line between the spheres' centers")
	t.Log("  will bounce between the spheres until the number of reflections is exceeded.")
	t.Log("  The resulting color should be black:")
	rayColor := scene.TestRay(ray)
	expected := color.Black

	assert.Equal(t, expected, rayColor)
}

func TestScene_TwoReflections(t *testing.T) {
	t.Log("Given a scene with a white skybox,")
	skyColor := color.White
	scene := render.NewScene(background.NewFlatColor(skyColor))

	sphereA := objects.NewSphere(core.NewVec3(0.0, 0.0, 0.0), 1)
	sphereB := objects.NewSphere(core.NewVec3(4.0, 3.0, 0.0), 1)
	sphereColor := color.New(0.2, 0.4, 0.6)
	randomizer := random.NewFakeRandomGenerator()
	material := materials.NewDiffusive(sphereColor, randomizer)
	scene.Add(sphereA, material)
	scene.Add(sphereB, material)
	t.Logf("and two spheres %v and %v with diffusive color %v,\n",
		sphereA, sphereB, sphereColor)

	firstHitPoint := core.NewVec3(math32.Sqrt(2)/2, math32.Sqrt(2)/2, 0.0)
	rayOrigin := core.NewVec3(10.0, 0.0, 0.0)
	ray := core.NewRay(rayOrigin, firstHitPoint.Sub(rayOrigin))
	t.Logf("  a ray %v should hit the first sphere at point %v,\n", ray, firstHitPoint)
	t.Log("  get reflected and hit the second sphere at point [3.0 3.0 0.0],")
	t.Log("  then get reflected towards the sky:")
	rayColor := scene.TestRay(ray)
	expected := sphereColor.MulColor(sphereColor)

	assert.Equal(t, expected, rayColor)
}
