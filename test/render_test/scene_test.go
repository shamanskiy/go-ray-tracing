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
	"github.com/stretchr/testify/assert"
)

var anyPoint = core.NewVec3(999, 666, 333)
var anyDirection = core.NewVec3(1, 2, 3)
var OBJECT_COLOR = color.Red
var OTHER_OBJECT_COLOR = color.Green
var BACKGROUND_COLOR = color.Blue
var randomizer = random.NewRandomGenerator()

func TestScene_ShouldReturnBackgroundColor_IfSceneEmpty(t *testing.T) {
	scene := emptyScene()
	ray := core.NewRay(anyPoint, anyDirection)

	rayColor := scene.TestRay(ray)

	assert.Equal(t, BACKGROUND_COLOR, rayColor)
}

func TestScene_ShouldReturnObjectColorMixedWithBackgroundColor_IfObjectHitOnce(t *testing.T) {
	scene := emptyScene()
	sphere := objects.NewSphere(core.NewVec3(0, 0, 0), 1)
	material := materials.NewDiffusive(OBJECT_COLOR, randomizer)
	scene.Add(sphere, material)
	ray := core.NewRay(core.NewVec3(2, 0, 0), core.NewVec3(-1, 0, 0))

	rayColor := scene.TestRay(ray)

	expectedColor := material.Color().MulColor(BACKGROUND_COLOR)
	assert.Equal(t, expectedColor, rayColor)
}

func TestScene_ShouldHitClosestObject(t *testing.T) {
	scene := emptyScene()
	sphere1 := objects.NewSphere(core.NewVec3(0, 0, 0), 1)
	material1 := materials.NewDiffusive(OBJECT_COLOR, randomizer)
	scene.Add(sphere1, material1)
	sphere2 := objects.NewSphere(core.NewVec3(-10, 0, 0), 1)
	material2 := materials.NewDiffusive(OTHER_OBJECT_COLOR, randomizer)
	scene.Add(sphere2, material2)
	ray := core.NewRay(core.NewVec3(2, 0, 0), core.NewVec3(-1, 0, 0))

	rayColor := scene.TestRay(ray)

	assert.Equal(t, material1.Color().MulColor(BACKGROUND_COLOR), rayColor)
}

func TestScene_ShouldReflectOfFirstObjectAndHitSecondObject(t *testing.T) {
	scene := reflectiveXYAngleScene()
	ray := core.NewRay(core.NewVec3(2, 1, 0), core.NewVec3(-1, -1, 0))

	rayColor := scene.TestRay(ray)

	expectedColor := OBJECT_COLOR.MulColor(OTHER_OBJECT_COLOR).MulColor(BACKGROUND_COLOR)
	assert.Equal(t, expectedColor, rayColor)
}

func TestScene_ShouldHitOnlySecondPlane_BecauseOfMiniminHitRayParameter(t *testing.T) {
	scene := reflectiveXYAngleScene()
	scene.SetMinRayHitParameter(2)
	ray := core.NewRay(core.NewVec3(2, 1, 0), core.NewVec3(-1, -1, 0))

	rayColor := scene.TestRay(ray)

	expectedColor := OTHER_OBJECT_COLOR.MulColor(BACKGROUND_COLOR)
	assert.Equal(t, expectedColor, rayColor)
}

func TestScene_ShouldColorRayBlack_IfMaxNumberOfReflectionsExceeded(t *testing.T) {
	scene := reflectiveXYAngleScene()
	scene.SetMaxReflectionDepth(1)
	ray := core.NewRay(core.NewVec3(2, 1, 0), core.NewVec3(-1, -1, 0))

	rayColor := scene.TestRay(ray)

	expectedColor := OBJECT_COLOR.MulColor(color.Black)
	assert.Equal(t, expectedColor, rayColor)
}

func emptyScene() *render.Scene {
	return render.NewScene(background.NewFlatColor(BACKGROUND_COLOR))
}

func reflectiveXYAngleScene() *render.Scene {
	scene := emptyScene()
	plane1 := objects.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	material1 := materials.NewReflective(OBJECT_COLOR, randomizer)
	scene.Add(plane1, material1)
	plane2 := objects.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	material2 := materials.NewReflective(OTHER_OBJECT_COLOR, randomizer)
	scene.Add(plane2, material2)
	return scene
}
