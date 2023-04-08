package main

import (
	"image/png"
	"os"
	"runtime"

	"github.com/Shamanskiy/go-ray-tracer/src/camera"
	"github.com/Shamanskiy/go-ray-tracer/src/camera/image"
	"github.com/Shamanskiy/go-ray-tracer/src/camera/log"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/background"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
)

var randomizer = random.NewRandomGenerator()

const SMALL_SPHERE_GRID_SIZE = 11

func main() {
	scene := makeScene()
	camera := makeCamera()
	image := camera.Render(scene)
	saveImage(image, "megaScene.png")
}

func makeScene() scene.Scene {
	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	scene := scene.New(background)

	floor := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	scene.Add(floor, materials.NewDiffusive(color.GrayMedium, randomizer))

	mattSphere := geometries.NewSphere(core.NewVec3(-2.5, 1.0, 1), 1)
	scene.Add(mattSphere, materials.NewDiffusive(color.Red, randomizer))
	glassSphere := geometries.NewSphere(core.NewVec3(0.0, 1.0, 0.15), 1)
	scene.Add(glassSphere, materials.NewTransparent(1.5, color.White, randomizer))
	mirrorSphere := geometries.NewSphere(core.NewVec3(2.5, 1.0, 0.0), 1)
	scene.Add(mirrorSphere, materials.NewReflectiveFuzzy(color.GrayLight, 0.03, randomizer))

	bigSpheres := []geometries.Sphere{mattSphere, glassSphere, mirrorSphere}
	makeGridOfRandomSpheres(scene, SMALL_SPHERE_GRID_SIZE, bigSpheres)

	return scene
}

func makeGridOfRandomSpheres(scene *scene.SceneImpl, gridSize int, bigSpheres []geometries.Sphere) {
	for i := -gridSize; i < gridSize; i++ {
		for j := -gridSize; j < gridSize; j++ {
			sphere := makeRandomSmallSphere(i, j)
			for tooCloseToBigSpheres(sphere, bigSpheres) {
				sphere = makeRandomSmallSphere(i, j)
			}

			material := makeRandomMaterial()
			scene.Add(sphere, material)
		}
	}
}

func makeRandomSmallSphere(i, j int) geometries.Sphere {
	center := core.NewVec3(
		core.Real(i)+0.8*randomizer.Real(),
		0.2,
		core.Real(j)+0.8*randomizer.Real(),
	)
	return geometries.NewSphere(center, 0.2)
}

func makeRandomMaterial() materials.Material {
	materialChoice := randomizer.Real()
	randomColor := color.FromVec3(randomizer.Vec3())
	if materialChoice < 0.7 {
		return materials.NewDiffusive(randomColor, randomizer)
	} else if materialChoice < 0.9 {
		metallicColor := color.New(1., 1., 1.).Add(randomColor).Mul(0.5)
		fuzziness := 0.1 * randomizer.Real()
		return materials.NewReflectiveFuzzy(metallicColor, fuzziness, randomizer)
	} else {
		return materials.NewTransparent(1.5, color.White, randomizer)
	}
}

func tooCloseToBigSpheres(sphere geometries.Sphere, bigSpheres []geometries.Sphere) bool {
	for _, bigSphere := range bigSpheres {
		if sphere.InContactWith(bigSphere) {
			return true
		}
	}
	return false
}

func makeCamera() *camera.Camera {
	settings := camera.CameraSettings{
		VerticalFOV:      75,
		AspectRatio:      16. / 9.,
		ImagePixelHeight: 360,
		LookFrom:         core.NewVec3(3.5, 1.35, 1.9),
		LookAt:           core.NewVec3(3., 1.25, 1.5),
		Antialiasing:     4,
		ProgressChan:     log.NewProgressBar(),
		NumRenderThreads: runtime.NumCPU(),
	}

	return camera.NewCamera(&settings, randomizer)
}

func saveImage(image *image.Image, filename string) {
	defer log.TimeExecution("save image")()
	file, _ := os.Create(filename)
	defer file.Close()
	rgbaImage := image.ConvertToRGBA()
	png.Encode(file, rgbaImage)
}
