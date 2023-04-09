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

const (
	MAX_RAY_REFLECTIONS = 10
	CAMERA_Z            = 1.2
)

func main() {
	scene := makeScene()
	camera := makeCamera()
	image := camera.Render(scene)
	saveImage(image, "cornellBox.png")
}

func makeScene() scene.Scene {
	background := background.NewFlatColor(color.White)
	scene := scene.New(background, scene.MaxRayReflections(MAX_RAY_REFLECTIONS))

	floor := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 1, 0))
	scene.Add(floor, materials.NewDiffusive(color.White, randomizer))

	topLight := geometries.NewPlane(core.NewVec3(0, 1, 0), core.NewVec3(0, 1, 0))
	scene.Add(topLight, materials.NewDiffusive(color.White, randomizer))

	backWall := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(0, 0, 1))
	scene.Add(backWall, materials.NewReflective(color.GrayMedium, randomizer))

	frontWall := geometries.NewPlane(core.NewVec3(0, 0, CAMERA_Z), core.NewVec3(0, 0, 1))
	scene.Add(frontWall, materials.NewReflective(color.GrayMedium, randomizer))

	leftWall := geometries.NewPlane(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	scene.Add(leftWall, materials.NewDiffusive(color.Red, randomizer))

	rightWall := geometries.NewPlane(core.NewVec3(1, 0, 0), core.NewVec3(1, 0, 0))
	scene.Add(rightWall, materials.NewDiffusive(color.Green, randomizer))

	sphere := geometries.NewSphere(core.NewVec3(0.5, 0.2, 0.5), 0.2)
	scene.Add(sphere, materials.NewDiffusiveLight(color.White))

	return scene
}

func makeCamera() *camera.Camera {
	settings := camera.CameraSettings{
		VerticalFOV:      90,
		AspectRatio:      1.,
		ImagePixelHeight: 360 * 2,
		LookFrom:         core.NewVec3(0.75, 0.5, CAMERA_Z),
		LookAt:           core.NewVec3(0.5, 0.5, 0.5),
		Antialiasing:     10,
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
