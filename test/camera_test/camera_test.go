package camera_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/camera"
	"github.com/Shamanskiy/go-ray-tracer/src/camera/image"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/background"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/stretchr/testify/assert"
)

var cameraSettings = camera.CameraSettings{
	VerticalFOV:      90,
	AspectRatio:      2,
	ImagePixelHeight: 10,
	LookAt:           core.NewVec3(0, 0, -1),
	Antialiasing:     1,
}

func TestCamera_ShouldColorAllPixelsSameColor_IfEmptySceneWithFlatColorBackground(t *testing.T) {
	randomizer := random.NewRandomGenerator()
	camera := camera.NewCamera(&cameraSettings, randomizer)
	scene := scene.New(background.NewFlatColor(color.Red))

	image := camera.Render(scene)

	assert.Equal(t, cameraSettings.ImagePixelHeight, image.Height())
	assert.Equal(t, cameraSettings.ImagePixelHeight*2, image.Width())
	assertAllPixelsColor(t, image, color.Red)
}

func TestCamera_ShouldColorCentralPixelWithObjectColor(t *testing.T) {
	randomizer := random.NewRandomGenerator()
	camera := camera.NewCamera(&cameraSettings, randomizer)
	scene := scene.New(background.NewFlatColor(color.White))
	scene.Add(geometries.NewSphere(core.NewVec3(0, 0, -2), 0.5), materials.NewDiffusive(color.Red, randomizer))

	image := camera.Render(scene)

	assertCornerPixelsColor(t, image, color.White)
	assert.Equal(t, color.Red, image.PixelColor(10, 5))
}

func assertAllPixelsColor(t *testing.T, image *image.Image, color color.Color) {
	for x := 0; x < image.Width(); x++ {
		for y := 0; y < image.Height(); y++ {
			assert.Equal(t, color, image.PixelColor(x, y))
		}
	}
}

func assertCornerPixelsColor(t *testing.T, image *image.Image, color color.Color) {
	width := image.Width()
	height := image.Height()
	assert.Equal(t, color, image.PixelColor(0, 0))
	assert.Equal(t, color, image.PixelColor(width-1, 0))
	assert.Equal(t, color, image.PixelColor(0, height-1))
	assert.Equal(t, color, image.PixelColor(width-1, height-1))
}
