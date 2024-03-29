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

const SMALL_SPHERE_GRID_SIZE = 11

func main() {
	scene := makeScene()
	camera := makeCamera()
	image := camera.Render(scene)
	image.SaveRGBAToPNG("megaScene.png")
}

func makeScene() scene.Scene {
	objects := []scene.Object{}

	floor := geometries.NewSphere(core.NewVec3(0, -500, 0), 500)
	floorMaterial := materials.NewDiffusive(color.GrayMedium, randomizer)
	objects = append(objects, scene.Object{Hittable: floor, Material: floorMaterial})

	sun := geometries.NewSphere(core.NewVec3(100, 200, 100), 50)
	sunMaterial := materials.NewDiffusiveLight(color.White, 1.)
	objects = append(objects, scene.Object{Hittable: sun, Material: sunMaterial})

	mattSphere := geometries.NewSphere(core.NewVec3(-2.5, 1.0, 1), 1)
	mattSphereMaterial := materials.NewDiffusive(color.Red, randomizer)
	objects = append(objects, scene.Object{Hittable: mattSphere, Material: mattSphereMaterial})

	glassSphere := geometries.NewSphere(core.NewVec3(0.0, 1.0, 0.15), 1)
	glassSphereMaterial := materials.NewTransparent(1.5, color.White, randomizer)
	objects = append(objects, scene.Object{Hittable: glassSphere, Material: glassSphereMaterial})

	mirrorSphere := geometries.NewSphere(core.NewVec3(2.5, 1.0, 0.0), 1)
	mirrorSphereMaterial := materials.NewReflectiveFuzzy(color.GrayLight, 0.02, randomizer)
	objects = append(objects, scene.Object{Hittable: mirrorSphere, Material: mirrorSphereMaterial})

	bigSpheres := []geometries.Sphere{mattSphere, glassSphere, mirrorSphere}
	objects = append(objects, makeGridOfRandomSpheres(SMALL_SPHERE_GRID_SIZE, bigSpheres)...)

	background := background.NewVerticalGradient(color.White, color.SkyBlue)
	return scene.New(objects, background)
}

func makeGridOfRandomSpheres(gridSize int, bigSpheres []geometries.Sphere) []scene.Object {
	objects := []scene.Object{}
	for i := -gridSize; i < gridSize; i++ {
		for j := -gridSize; j < gridSize; j++ {
			sphere := makeRandomSmallSphere(i, j)
			for tooCloseToBigSpheres(sphere, bigSpheres) {
				sphere = makeRandomSmallSphere(i, j)
			}

			material := makeRandomMaterial()
			objects = append(objects, scene.Object{Hittable: sphere, Material: material})
		}
	}
	return objects
}

func makeRandomSmallSphere(i, j int) geometries.Sphere {
	center := core.NewVec3(
		core.Real(i)+0.8*randomizer.Real(),
		0.2,
		core.Real(j)+0.8*randomizer.Real(),
	)
	return geometries.NewSphere(center, 0.2)
}

func makeRandomMaterial() materials.Material {
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
		VerticalFOV:         75,
		AspectRatio:         16. / 9.,
		ImagePixelHeight:    360 * 3,
		LookFrom:            core.NewVec3(3.5, 1.35, 1.9),
		LookAt:              core.NewVec3(3., 1.25, 1.5),
		Antialiasing:        20,
		ProgressChan:        log.NewProgressBar(),
		NumRenderThreads:    runtime.NumCPU(),
		DefocusBlurStrength: 0.005,
	}

	return camera.NewCamera(&settings, randomizer)
}
