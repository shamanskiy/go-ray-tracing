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

const (
	MAX_RAY_REFLECTIONS = 10
	CAMERA_Z            = 1.2
)

func main() {
	scene := makeScene()
	camera := makeCamera()
	image := camera.Render(scene)
	image.SaveRGBAToPNG("cornellBox.png")
}

func makeScene() scene.Scene {
	objects := []scene.Object{}

	floor := geometries.NewQuad(
		core.NewVec3(0, 0, 0),
		core.NewVec3(0, 0, 1),
		core.NewVec3(1, 0, 1),
		core.NewVec3(1, 0, 0))
	floorMaterial := materials.NewDiffusive(color.White, randomizer)
	objects = append(objects, scene.Object{Hittable: floor, Material: floorMaterial})

	topLight := geometries.NewQuad(
		core.NewVec3(0, 1, 0),
		core.NewVec3(1, 1, 0),
		core.NewVec3(1, 1, 1),
		core.NewVec3(0, 1, 1))
	topLightMaterial := materials.NewDiffusive(color.White, randomizer)
	objects = append(objects, scene.Object{Hittable: topLight, Material: topLightMaterial})

	backWall := geometries.NewQuad(
		core.NewVec3(0, 0, 0),
		core.NewVec3(1, 0, 0),
		core.NewVec3(1, 1, 0),
		core.NewVec3(0, 1, 0))
	backWallMaterial := materials.NewReflective(color.GrayMedium, randomizer)
	objects = append(objects, scene.Object{Hittable: backWall, Material: backWallMaterial})

	leftWall := geometries.NewQuad(
		core.NewVec3(0, 0, 0),
		core.NewVec3(0, 1, 0),
		core.NewVec3(0, 1, 1),
		core.NewVec3(0, 0, 1))
	leftWallMaterial := materials.NewDiffusive(color.Red, randomizer)
	objects = append(objects, scene.Object{Hittable: leftWall, Material: leftWallMaterial})

	rightWall := geometries.NewQuad(
		core.NewVec3(1, 0, 0),
		core.NewVec3(1, 0, 1),
		core.NewVec3(1, 1, 1),
		core.NewVec3(1, 1, 0))
	rightWallMaterial := materials.NewDiffusive(color.Green, randomizer)
	objects = append(objects, scene.Object{Hittable: rightWall, Material: rightWallMaterial})

	sphere := geometries.NewSphere(core.NewVec3(0.5, 0.2, 0.5), 0.2)
	sphereMaterial := materials.NewDiffusiveLight(color.White)
	objects = append(objects, scene.Object{Hittable: sphere, Material: sphereMaterial})

	background := background.NewFlatColor(color.White)
	return scene.New(objects, background, scene.MaxRayReflections(MAX_RAY_REFLECTIONS))
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
