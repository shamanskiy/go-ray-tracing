package main

import (
	"runtime"

	"github.com/Shamanskiy/go-ray-tracer/src/camera"
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
	image.SaveRGBAToPNG("threeSpheres.png")
}

func makeScene() scene.Scene {
	objects := []scene.Object{}

	floor := geometries.NewSphere(core.NewVec3(0, -500.5, 0), 500)
	floorMaterial := materials.NewDiffusive(color.GrayMedium, randomizer)
	objects = append(objects, scene.Object{Hittable: floor, Material: floorMaterial})

	mattSphere := geometries.NewSphere(core.NewVec3(0, 0, -1), 0.5)
	mattSphereMaterial := materials.NewDiffusive(color.Red, randomizer)
	objects = append(objects, scene.Object{Hittable: mattSphere, Material: mattSphereMaterial})

	smallMattSphere := geometries.NewSphere(core.NewVec3(0.25, -0.4, -0.5), 0.1)
	smallMattSphereMaterial := materials.NewDiffusive(color.SkyBlue, randomizer)
	objects = append(objects, scene.Object{Hittable: smallMattSphere, Material: smallMattSphereMaterial})

	mirrowSphere := geometries.NewSphere(core.NewVec3(1, 0, -1), 0.5)
	mirrowSphereMaterial := materials.NewReflectiveFuzzy(color.GrayLight, 0.05, randomizer)
	objects = append(objects, scene.Object{Hittable: mirrowSphere, Material: mirrowSphereMaterial})

	glassShellOuter := geometries.NewSphere(core.NewVec3(-1, 0, -1), 0.5)
	glassShellInner := geometries.NewSphere(core.NewVec3(-1, 0, -1), -0.4)
	glassMaterial := materials.NewTransparent(1.5, color.White, randomizer)
	objects = append(objects, scene.Object{Hittable: glassShellOuter, Material: glassMaterial})
	objects = append(objects, scene.Object{Hittable: glassShellInner, Material: glassMaterial})

	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	return scene.New(objects, background)
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
