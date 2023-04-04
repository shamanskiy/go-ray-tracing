// This is a demo app that renders three spheres with
// different materials.
package main

import (
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

func makeScene() scene.Scene {
	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	scene := scene.New(background)

	// matt green ball
	scene.Add(geometries.NewSphere(core.NewVec3(0, 0, -1), 0.5),
		materials.NewDiffusive(color.Red, randomizer))

	// mirrow ball
	scene.Add(geometries.NewSphere(core.NewVec3(1, 0, -1), 0.5),
		materials.NewReflectiveFuzzy(color.Golden, 0.1, randomizer))

	// glass shell
	scene.Add(geometries.NewSphere(core.NewVec3(-1, 0, -1), 0.5),
		materials.NewTransparent(1.5, color.White, randomizer))
	scene.Add(geometries.NewSphere(core.NewVec3(-1, 0, -1), -0.4),
		materials.NewTransparent(1.5, color.White, randomizer))

	// floor
	scene.Add(geometries.NewPlane(core.NewVec3(0, -0.5, 0), core.NewVec3(0, 1, 0)),
		materials.NewDiffusive(color.GrayMedium, randomizer))
	return scene
}

func makeCamera() *camera.Camera {
	settings := camera.DefaultCameraSettings()
	settings.AspectRatio = 16. / 9.
	settings.ImagePixelHeight = 360
	settings.Antialiasing = 4
	settings.LookFrom = core.NewVec3(0, 0, 0.15)
	settings.LookAt = core.NewVec3(0, 0, -1)

	settings.ProgressChan = log.ProgressBar(100, "rendering")

	camera := camera.NewCamera(&settings, randomizer)

	return camera
}

func saveImage(image *image.Image, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()
	image.SaveAsPNG(file)
}

func main() {
	scene := makeScene()
	camera := makeCamera()
	image := camera.Render(scene)
	saveImage(image, "threeSpheres.png")
}
