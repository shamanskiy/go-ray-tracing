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

var pixelwiseTests = []string{
	"redDiffusiveSphere",
	"grayReflectiveSphere",
	"redDiffusiveTriangle",
}

func TestPixelwiseImageComparison(t *testing.T) {
	camera := makeCamera()

	for _, testName := range pixelwiseTests {
		t.Run(testName, func(t *testing.T) {
			scene := makeScene(testName)

			img := camera.Render(scene)

			assertImage(t, testName, img)
		})
	}
}

func makeCamera() *camera.Camera {
	settings := camera.CameraSettings{
		VerticalFOV:      70,
		AspectRatio:      16. / 9.,
		ImagePixelHeight: 360,
		LookFrom:         core.NewVec3(0, 0, 1),
		LookAt:           core.NewVec3(0, 0, 0),
		Antialiasing:     4,
		NumRenderThreads: runtime.NumCPU(),
	}

	return camera.NewCamera(&settings, randomizer)
}

func assertImage(t *testing.T, testName string, renderedImage *image.Image) {
	benchmarkImageName := testName + ".png"
	benchmarkImageFile, err := fs.Open(benchmarkImageName)
	test.PanicOnErr(err)
	benchmarkImage, err := png.Decode(benchmarkImageFile)
	test.PanicOnErr(err)

	assert.Equal(t, benchmarkImage, renderedImage.ConvertToRGBA())
}

func makeScene(testName string) scene.Scene {
	switch testName {
	case "redDiffusiveSphere":
		return redDiffusiveSphereScene()
	case "grayReflectiveSphere":
		return grayReflectiveSphereScene()
	case "redDiffusiveTriangle":
		return redDiffusiveTriangleScene()
	default:
		panic("unknown testName: " + testName)
	}
}

func redDiffusiveSphereScene() scene.Scene {
	objects := []scene.Object{}

	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 0.5)
	material := materials.NewDiffusive(color.Red, randomizer)
	objects = append(objects, scene.Object{Hittable: sphere, Material: material})

	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	return scene.New(objects, background)
}

func grayReflectiveSphereScene() scene.Scene {
	objects := []scene.Object{}

	sphere := geometries.NewSphere(core.NewVec3(0, 0, 0), 0.5)
	material := materials.NewReflective(color.GrayLight, randomizer)
	objects = append(objects, scene.Object{Hittable: sphere, Material: material})

	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	return scene.New(objects, background)
}

func redDiffusiveTriangleScene() scene.Scene {
	objects := []scene.Object{}

	triangle := geometries.NewTriangleWithNormals(
		core.NewVec3(-0.5, -0.5, 0),
		core.NewVec3(0.5, -0.5, 0),
		core.NewVec3(0, 0.5, 0),
		core.NewVec3(0, 0, 1),
		core.NewVec3(0, 0, 1),
		core.NewVec3(0, 0, 1))
	material := materials.NewDiffusive(color.Red, randomizer)
	objects = append(objects, scene.Object{Hittable: triangle, Material: material})

	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	return scene.New(objects, background)
}
