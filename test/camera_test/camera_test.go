package camera_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/camera"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/background"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

func TestCamera_Default(t *testing.T) {
	t.Log("Default camera without randomness")
	settings := camera.DefaultCameraSettings()
	randomizer := random.NewFakeRandomGenerator()

	camera := camera.NewCamera(&settings, randomizer)

	assert.Equal(t, settings.LookFrom, camera.Origin)

	ray := camera.GetRay(0.5, 0.5)
	test.AssertInDeltaVec3(t, settings.LookFrom, ray.Origin(), core.Tolerance)
	test.AssertInDeltaVec3(t, settings.LookAt, ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0.5, 0)
	test.AssertInDeltaVec3(t, core.NewVec3(0, 1, -1), ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0.5, 1)
	test.AssertInDeltaVec3(t, core.NewVec3(0, -1, -1), ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0, 0)
	test.AssertInDeltaVec3(t, core.NewVec3(-2, 1, -1), ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(1, 1)
	test.AssertInDeltaVec3(t, core.NewVec3(2, -1, -1), ray.Eval(1), core.Tolerance)
}

func TestCamera_Custom(t *testing.T) {
	t.Log("Camera with custom settings without randomness")
	settings := camera.DefaultCameraSettings()
	settings.LookAt = core.NewVec3(3, 0, 4)
	settings.AspectRatio = 1
	randomizer := random.NewFakeRandomGenerator()

	camera := camera.NewCamera(&settings, randomizer)

	test.AssertInDeltaVec3(t, settings.LookFrom, camera.Origin, core.Tolerance)

	ray := camera.GetRay(0.5, 0.5)
	test.AssertInDeltaVec3(t, settings.LookFrom, ray.Origin(), core.Tolerance)
	test.AssertInDeltaVec3(t, settings.LookAt, ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0.5, 0)
	test.AssertInDeltaVec3(t, core.NewVec3(3, 5, 4), ray.Eval(1), core.Tolerance)

	ray = camera.GetRay(0, 0.5)
	test.AssertInDeltaVec3(t, core.NewVec3(7, 0, 1), ray.Eval(1), core.Tolerance)
}

func TestCamera_indexToU(t *testing.T) {
	t.Log("Camera with 100 px height and 1:1 aspect ratio")
	settings := camera.DefaultCameraSettings()
	settings.ImagePixelHeight = 100
	settings.AspectRatio = 2.0
	randomizer := random.NewFakeRandomGenerator()

	camera := camera.NewCamera(&settings, randomizer)

	assert.Equal(t, 100, camera.PixelHeight)
	assert.Equal(t, 200, camera.PixelWidth)

	assert.EqualValues(t, 0, camera.IndexToU(0))
	assert.EqualValues(t, 0.5, camera.IndexToU(100))
	assert.EqualValues(t, 1, camera.IndexToU(200))

	assert.EqualValues(t, 0, camera.IndexToV(0))
	assert.EqualValues(t, 0.5, camera.IndexToV(50))
	assert.EqualValues(t, 1, camera.IndexToV(100))
}

func TestCamera_RenderEmptyScene(t *testing.T) {
	t.Log("Given an empty scene with white background")
	scene := scene.New(background.NewFlatColor(color.White))

	imageSize := 2
	t.Logf("and a camera with %vx%v resolution,\n", imageSize, imageSize)
	settings := camera.DefaultCameraSettings()
	settings.ImagePixelHeight = imageSize
	settings.AspectRatio = 1
	settings.Antialiasing = 1
	randomizer := random.NewFakeRandomGenerator()
	camera := camera.NewCamera(&settings, randomizer)

	t.Logf("  the rendered image should be a %vx%v white square:\n", imageSize, imageSize)
	camera.Render(scene)

	// assert.Equal(t, imageSize, renderedImage.Bounds().Size().X)
	// assert.Equal(t, imageSize, renderedImage.Bounds().Size().Y)

	// for x := 0; x < imageSize; x++ {
	// 	for y := 0; y < imageSize; y++ {
	// 		assert.Equal(t, color.White, renderedImage.PixelColor(x, y))
	// 	}
	// }

}
