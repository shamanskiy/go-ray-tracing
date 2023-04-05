// This is a big scene with three big central spheres and
// 450+ small randomly generated spheres.
// ATTENTION: the scene can take a while to render
package main

import (
	"image/png"
	"os"

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

func main() {
	scene := makeScene()
	camera := makeCamera()
	image := camera.Render(scene)
	saveImage(image, "megaScene.png")
}

func makeScene() scene.Scene {
	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	scene := scene.New(background)

	// Huge sphere = floor
	scene.Add(geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0)),
		materials.NewDiffusive(color.GrayMedium, randomizer))

	// Three main spheres
	sphereRed := geometries.NewSphere(core.NewVec3(-2.5, 1.0, 1), 1)
	scene.Add(sphereRed, materials.NewDiffusive(color.Red, randomizer))
	sphereGlass := geometries.NewSphere(core.NewVec3(0.0, 1.0, 0.15), 1)
	scene.Add(sphereGlass, materials.NewTransparent(1.5, color.White, randomizer))
	sphereMirror := geometries.NewSphere(core.NewVec3(2.5, 1.0, 0.0), 1)
	scene.Add(sphereMirror, materials.NewReflectiveFuzzy(color.GrayLight, 0.03, randomizer))

	// Little spheres randomly offset from a regular grid
	gridSize := 11
	bigSpheres := []geometries.Sphere{sphereRed, sphereGlass, sphereMirror}
	for i := -gridSize; i < gridSize; i++ {
		for j := -gridSize; j < gridSize; j++ {
			sphere := genRandomSmallSphere(i, j)
			for tooCloseToBigSpheres(sphere, bigSpheres) {
				sphere = genRandomSmallSphere(i, j)
			}

			material := genRandomMaterial()
			scene.Add(sphere, material)
		}
	}

	return scene
}

func genRandomSmallSphere(i, j int) geometries.Sphere {
	center := core.NewVec3(
		core.Real(i)+0.8*randomizer.Real(),
		0.2,
		core.Real(j)+0.8*randomizer.Real(),
	)
	return geometries.NewSphere(center, 0.2)
}

func genRandomMaterial() materials.Material {
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
		Antialiasing:     10,
		ProgressChan:     log.ProgressBar(100, "rendering"),
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
