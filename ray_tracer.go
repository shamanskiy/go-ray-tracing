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
		materials.Reflective{Color: core.GrayLight})

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
	camera := render.Camera{
		Origin:            core.Vec3{0.0, 0.0, 0.0},
		Upper_left_corner: core.Vec3{-2.0, 1.0, -1.0},
		Horizontal:        core.Vec3{4.0, 0.0, 0.0},
		Vertical:          core.Vec3{0.0, -2.0, 0.0},
	}

	return &camera
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

	saveImage(img, "image.png")
}
