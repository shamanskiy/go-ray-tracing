// This is a demo app that renders three spheres with
// different materials.
package main

import (
	"image"
	"image/png"
	"os"

	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/render"
)

var randomizer = random.NewRandomGenerator()

func makeScene() *render.Scene {
	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	scene := render.NewScene(background)

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

func makeCamera() *render.Camera {
	settings := render.DefaultCameraSettings()
	settings.AspectRatio = 16. / 9.
	settings.ImagePixelHeight = 360
	settings.Antialiasing = 4
	settings.LookFrom = core.NewVec3(0, 0, 0.15)
	settings.LookAt = core.NewVec3(0, 0, -1)

	camera := render.NewCamera(&settings, randomizer)

	return camera
}

func saveImage(img *image.RGBA, filename string) {
	f, _ := os.Create(filename)
	defer f.Close()
	png.Encode(f, img)
}

func main() {
	scene := makeScene()
	camera := makeCamera()
	img := camera.Render(scene)
	saveImage(img, "threeSpheres.png")
}
