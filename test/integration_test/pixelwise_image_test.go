package integration_test

import (
	"image/png"
	"runtime"
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
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

var randomizer = random.NewFakeRandomGenerator()

func TestPixelwiseImageComparison(t *testing.T) {
	scene := makeScene()
	camera := makeCamera()

	img := camera.Render(scene)

	assertImage(t, "singleRedSphere.png", img)
}

func makeScene() scene.Scene {
	objects := []scene.Object{}

	mattSphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 0.5)
	mattSphereMaterial := materials.NewDiffusive(color.Red, randomizer)
	objects = append(objects, scene.Object{Hittable: mattSphere, Material: mattSphereMaterial})

	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	return scene.New(objects, background)
}

func makeCamera() *camera.Camera {
	settings := camera.CameraSettings{
		VerticalFOV:      70,
		AspectRatio:      16. / 9.,
		ImagePixelHeight: 360,
		LookFrom:         core.NewVec3(1, 1, 1),
		LookAt:           core.NewVec3(0, 0, 0),
		Antialiasing:     4,
		NumRenderThreads: runtime.NumCPU(),
	}

	return camera.NewCamera(&settings, randomizer)
}

func assertImage(t *testing.T, benchmarkImageName string, renderedImage *image.Image) {
	benchmarkImageFile, err := fs.Open(benchmarkImageName)
	test.PanicOnErr(err)
	benchmarkImage, err := png.Decode(benchmarkImageFile)
	test.PanicOnErr(err)

	assert.Equal(t, benchmarkImage, renderedImage.ConvertToRGBA())
}
