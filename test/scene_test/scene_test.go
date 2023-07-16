package scene_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/background"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/stretchr/testify/assert"
)

var anyPoint = core.NewVec3(999, 666, 333)
var anyDirection = core.NewVec3(1, 2, 3)
var OBJECT_COLOR = color.Red
var OTHER_OBJECT_COLOR = color.Green
var BACKGROUND_COLOR = color.Blue
var randomizer = random.NewRandomGenerator()
var noObjects = []scene.Object{}

func TestScene_ShouldReturnBackgroundColor_IfSceneEmpty(t *testing.T) {
	scene := scene.New(noObjects, flatBackground())
	ray := core.NewRay(anyPoint, anyDirection)

	rayColor := scene.TestRay(ray)

	assert.Equal(t, BACKGROUND_COLOR, rayColor)
}

func TestScene_ShouldReturnObjectColorMixedWithBackgroundColor_IfObjectHitOnce(t *testing.T) {
	scene := scene.New([]scene.Object{unitSphere(OBJECT_COLOR)}, flatBackground())
	ray := core.NewRay(core.NewVec3(2, 0, 0), core.NewVec3(-1, 0, 0))

	rayColor := scene.TestRay(ray)

	expectedColor := OBJECT_COLOR.MulColor(BACKGROUND_COLOR)
	assert.Equal(t, expectedColor, rayColor)
}

func TestScene_ShouldHitClosestObject(t *testing.T) {
	objects := []scene.Object{unitSphere(OBJECT_COLOR), unitSphere(OTHER_OBJECT_COLOR, core.NewVec3(-10, 0, 0))}
	scene := scene.New(objects, flatBackground())
	ray := core.NewRay(core.NewVec3(2, 0, 0), core.NewVec3(-1, 0, 0))

	rayColor := scene.TestRay(ray)

	assert.Equal(t, OBJECT_COLOR.MulColor(BACKGROUND_COLOR), rayColor)
}

func TestScene_ShouldReflectOfFirstObjectAndHitSecondObject(t *testing.T) {
	scene := scene.New(reflectiveXYAngle(), flatBackground())
	ray := core.NewRay(core.NewVec3(2, 1, 0), core.NewVec3(-1, -1, 0))

	rayColor := scene.TestRay(ray)

	expectedColor := OBJECT_COLOR.MulColor(OTHER_OBJECT_COLOR).MulColor(BACKGROUND_COLOR)
	assert.Equal(t, expectedColor, rayColor)
}

func TestScene_ShouldHitOnlySecondPlane_BecauseOfMiniminHitRayParameter(t *testing.T) {
	scene := scene.New(reflectiveXYAngle(), flatBackground(), scene.MinRayHitParameter(2))
	ray := core.NewRay(core.NewVec3(2, 1, 0), core.NewVec3(-1, -1, 0))

	rayColor := scene.TestRay(ray)

	expectedColor := OTHER_OBJECT_COLOR.MulColor(BACKGROUND_COLOR)
	assert.Equal(t, expectedColor, rayColor)
}

func TestScene_ShouldColorRayBlack_IfMaxNumberOfReflectionsExceeded(t *testing.T) {
	scene := scene.New(reflectiveXYAngle(), flatBackground(), scene.MaxRayReflections(1))
	ray := core.NewRay(core.NewVec3(2, 1, 0), core.NewVec3(-1, -1, 0))

	rayColor := scene.TestRay(ray)

	expectedColor := OBJECT_COLOR.MulColor(color.Black)
	assert.Equal(t, expectedColor, rayColor)
}

func flatBackground() background.Background {
	return background.NewFlatColor(BACKGROUND_COLOR)
}

func reflectiveXYAngle() []scene.Object {
	plane1 := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	material1 := materials.NewReflective(OBJECT_COLOR, randomizer)
	plane2 := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	material2 := materials.NewReflective(OTHER_OBJECT_COLOR, randomizer)

	return []scene.Object{
		{
			Hittable: plane1,
			Material: material1,
		},
		{
			Hittable: plane2,
			Material: material2,
		},
	}
}

func unitSphere(materialColor color.Color, translation ...core.Vec3) scene.Object {
	center := core.NewVec3(0, 0, 0)
	for _, vec := range translation {
		center = center.Add(vec)
	}

	sphere := geometries.NewSphere(center, 1)
	material := materials.NewDiffusive(materialColor, randomizer)
	return scene.Object{Hittable: sphere, Material: material}
}
