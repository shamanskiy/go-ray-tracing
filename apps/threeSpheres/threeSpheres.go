// This is a demo app that renders three spheres with
// different materials.
package main

import (
	"image"
	"image/png"
	"os"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/materials"
	"github.com/Shamanskiy/go-ray-tracer/objects"
	"github.com/Shamanskiy/go-ray-tracer/render"
)

func makeScene() *render.Scene {
	scene := render.Scene{SkyColorTop: core.SkyBlue, SkyColorBottom: core.White}

	// matt green ball
	scene.Add(objects.Sphere{Center: core.Vec3{0.0, 0.0, -1.0}, Radius: 0.5},
		materials.Diffusive{core.Green})

	// mirrow ball
	scene.Add(objects.Sphere{Center: core.Vec3{1.0, 0.0, -1.0}, Radius: 0.5},
		materials.NewReflectiveFuzzy(core.Golden, 0.1))

	// glass shell
	scene.Add(objects.Sphere{Center: core.Vec3{-1.0, 0.0, -1.0}, Radius: 0.5},
		materials.NewTransparent(1.5))
	scene.Add(objects.Sphere{Center: core.Vec3{-1.0, 0.0, -1.0}, Radius: -0.4},
		materials.NewTransparent(1.5))

	// Huge sphere = floor
	scene.Add(objects.Sphere{Center: core.Vec3{0.0, -100.5, -1.0}, Radius: 100.0},
		materials.Diffusive{core.GrayMedium})

	return &scene
}

func makeCamera() *render.Camera {
	settings := render.DefaultCameraSettings()
	settings.AspectRatio = 16. / 9.
	settings.ImagePixelHeight = 360
	settings.Antialiasing = 4
	settings.LookFrom = core.Vec3{0., 0., 0.15}
	settings.LookAt = core.Vec3{0., 0., -1}

	camera := render.NewCamera(&settings)

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
