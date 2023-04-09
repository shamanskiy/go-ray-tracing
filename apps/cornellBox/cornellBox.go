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

func main() {
	scene := makeScene()
	camera := makeCamera()
	image := camera.Render(scene)
	saveImage(image, "cornellBox.png")
}

func makeScene() scene.Scene {
	background := background.NewFlatColor(color.White)
	scene := scene.New(background)

	bottom := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	scene.Add(bottom, materials.NewDiffusive(color.White, randomizer))

	// top := geometries.NewPlane(core.NewVec3(0, 1, 0), core.NewVec3(0, 1, 0))
	// scene.Add(top, materials.NewDiffusive(color.White, randomizer))

	back := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 0, 1))
	scene.Add(back, materials.NewDiffusive(color.White, randomizer))

	left := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	scene.Add(left, materials.NewDiffusive(color.Red, randomizer))

	// right := geometries.NewPlane(core.NewVec3(1, 0, 0), core.NewVec3(1, 0, 0))
	// scene.Add(right, materials.NewDiffusive(color.Green, randomizer))

	return scene
}

func makeCamera() *camera.Camera {
	settings := camera.CameraSettings{
		VerticalFOV:      90,
		AspectRatio:      1.,
		ImagePixelHeight: 360,
		LookFrom:         core.NewVec3(0.5, 0.5, 1),
		LookAt:           core.NewVec3(0.5, 0.5, 0.5),
		Antialiasing:     1,
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
