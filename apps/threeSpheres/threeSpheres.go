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
	saveImage(image, "threeSpheres.png")
}

func makeScene() scene.Scene {
	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	scene := scene.New(background)

	floor := geometries.NewPlane(core.NewVec3(0, -0.5, 0), core.NewVec3(0, 1, 0))
	scene.Add(floor, materials.NewDiffusive(color.GrayMedium, randomizer))

	mattSphere := geometries.NewSphere(core.NewVec3(0, 0, -1), 0.5)
	scene.Add(mattSphere, materials.NewDiffusive(color.Red, randomizer))

	smallMattSphere := geometries.NewSphere(core.NewVec3(0.25, -0.4, -0.5), 0.1)
	scene.Add(smallMattSphere, materials.NewDiffusive(color.SkyBlue, randomizer))

	mirrowSphere := geometries.NewSphere(core.NewVec3(1, 0, -1), 0.5)
	scene.Add(mirrowSphere, materials.NewReflectiveFuzzy(color.GrayLight, 0.05, randomizer))

	glassShellOuter := geometries.NewSphere(core.NewVec3(-1, 0, -1), 0.5)
	glassShellInner := geometries.NewSphere(core.NewVec3(-1, 0, -1), -0.4)
	glassMaterial := materials.NewTransparent(1.5, color.White, randomizer)
	scene.Add(glassShellOuter, glassMaterial)
	scene.Add(glassShellInner, glassMaterial)

	scene.BuildBVH()
	return scene
}

func makeCamera() *camera.Camera {
	settings := camera.CameraSettings{
		VerticalFOV:         70,
		AspectRatio:         16. / 9.,
		ImagePixelHeight:    360 * 5,
		LookFrom:            core.NewVec3(0, 0, 0.5),
		LookAt:              core.NewVec3(0, 0, -1),
		Antialiasing:        4,
		ProgressChan:        log.NewProgressBar(),
		NumRenderThreads:    runtime.NumCPU(),
		DefocusBlurStrength: 0.05,
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
