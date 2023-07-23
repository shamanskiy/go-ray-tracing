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
	camera := makeCamera()
	scene := makeScene()
	image := camera.Render(scene)
	image.SaveRGBAToPNG("cornellBox.png")
}

func makeCamera() *camera.Camera {
	settings := camera.CameraSettings{
		VerticalFOV:      37.5,
		AspectRatio:      1.,
		ImagePixelHeight: 360,
		LookFrom:         core.NewVec3(278, 273, -800),
		LookAt:           core.NewVec3(278, 273, 0),
		Antialiasing:     40,
		ProgressChan:     log.NewProgressBar(),
		NumRenderThreads: runtime.NumCPU(),
	}

	return camera.NewCamera(&settings, randomizer)
}

func makeScene() scene.Scene {
	// http://www.graphics.cornell.edu/online/box/data.html
	objects := []scene.Object{}

	objects = append(objects, makeBox()...)
	objects = append(objects, makeTopLight())
	objects = append(objects, makeShortBlock()...)
	objects = append(objects, makeTallBlock()...)

	background := background.NewFlatColor(color.Black)
	return scene.New(objects, background)
}

func makeBox() []scene.Object {
	objects := []scene.Object{}

	floor := geometries.NewQuad(
		core.NewVec3(556.0, 0.0, 0.0),
		core.NewVec3(0.0, 0.0, 0.0),
		core.NewVec3(0.0, 0.0, 559.2),
		core.NewVec3(556.0, 0.0, 559.2))
	floorMaterial := materials.NewDiffusive(color.White, randomizer)
	objects = append(objects, scene.Object{Hittable: floor, Material: floorMaterial})

	ceiling := geometries.NewQuad(
		core.NewVec3(556.0, 548.8, 0.0),
		core.NewVec3(556.0, 548.8, 559.2),
		core.NewVec3(0.0, 548.8, 559.2),
		core.NewVec3(0.0, 548.8, 0.0))
	ceilingMaterial := materials.NewDiffusive(color.White, randomizer)
	objects = append(objects, scene.Object{Hittable: ceiling, Material: ceilingMaterial})

	backWall := geometries.NewQuad(
		core.NewVec3(556.0, 0.0, 559.2),
		core.NewVec3(0.0, 0.0, 559.2),
		core.NewVec3(0.0, 548.8, 559.2),
		core.NewVec3(556.0, 548.8, 559.2))
	backWallMaterial := materials.NewDiffusive(color.White, randomizer)
	objects = append(objects, scene.Object{Hittable: backWall, Material: backWallMaterial})

	leftWall := geometries.NewQuad(
		core.NewVec3(556.0, 0.0, 0.0),
		core.NewVec3(556.0, 0.0, 559.2),
		core.NewVec3(556.0, 548.8, 559.2),
		core.NewVec3(556.0, 548.8, 0.0))
	leftWallMaterial := materials.NewDiffusive(color.Red, randomizer)
	objects = append(objects, scene.Object{Hittable: leftWall, Material: leftWallMaterial})

	rightWall := geometries.NewQuad(
		core.NewVec3(0.0, 0.0, 559.2),
		core.NewVec3(0.0, 0.0, 0.0),
		core.NewVec3(0.0, 548.8, 0.0),
		core.NewVec3(0.0, 548.8, 559.2))
	rightWallMaterial := materials.NewDiffusive(color.Green, randomizer)
	objects = append(objects, scene.Object{Hittable: rightWall, Material: rightWallMaterial})

	return objects
}

func makeTopLight() scene.Object {
	topLight := geometries.NewQuad(
		core.NewVec3(343.0, 548.8, 227.0),
		core.NewVec3(343.0, 548.8, 332.0),
		core.NewVec3(213.0, 548.8, 332.0),
		core.NewVec3(213.0, 548.8, 227.0))
	topLightMaterial := materials.NewDiffusiveLight(color.White, 10.)
	return scene.Object{Hittable: topLight, Material: topLightMaterial}
}

func makeShortBlock() []scene.Object {
	objects := []scene.Object{}

	material := materials.NewDiffusive(color.White, randomizer)

	top := geometries.NewQuad(
		core.NewVec3(130.0, 165.0, 65.0),
		core.NewVec3(82.0, 165.0, 225.0),
		core.NewVec3(240.0, 165.0, 272.0),
		core.NewVec3(290.0, 165.0, 114.0))
	objects = append(objects, scene.Object{Hittable: top, Material: material})

	side1 := geometries.NewQuad(
		core.NewVec3(290.0, 0.0, 114.0),
		core.NewVec3(290.0, 165.0, 114.0),
		core.NewVec3(240.0, 165.0, 272.0),
		core.NewVec3(240.0, 0.0, 272.0))
	objects = append(objects, scene.Object{Hittable: side1, Material: material})

	side2 := geometries.NewQuad(
		core.NewVec3(130.0, 0.0, 65.0),
		core.NewVec3(130.0, 165.0, 65.0),
		core.NewVec3(290.0, 165.0, 114.0),
		core.NewVec3(290.0, 0.0, 114.0))
	objects = append(objects, scene.Object{Hittable: side2, Material: material})

	side3 := geometries.NewQuad(
		core.NewVec3(82.0, 0.0, 225.0),
		core.NewVec3(82.0, 165.0, 225.0),
		core.NewVec3(130.0, 165.0, 65.0),
		core.NewVec3(130.0, 0.0, 65.0))
	objects = append(objects, scene.Object{Hittable: side3, Material: material})

	side4 := geometries.NewQuad(
		core.NewVec3(240.0, 0.0, 272.0),
		core.NewVec3(240.0, 165.0, 272.0),
		core.NewVec3(82.0, 165.0, 225.0),
		core.NewVec3(82.0, 0.0, 225.0))
	objects = append(objects, scene.Object{Hittable: side4, Material: material})

	return objects
}

func makeTallBlock() []scene.Object {
	objects := []scene.Object{}

	material := materials.NewDiffusive(color.White, randomizer)

	top := geometries.NewQuad(
		core.NewVec3(423.0, 330.0, 247.0),
		core.NewVec3(265.0, 330.0, 296.0),
		core.NewVec3(314.0, 330.0, 456.0),
		core.NewVec3(472.0, 330.0, 406.0))
	objects = append(objects, scene.Object{Hittable: top, Material: material})

	side1 := geometries.NewQuad(
		core.NewVec3(423.0, 0.0, 247.0),
		core.NewVec3(423.0, 330.0, 247.0),
		core.NewVec3(472.0, 330.0, 406.0),
		core.NewVec3(472.0, 0.0, 406.0))
	objects = append(objects, scene.Object{Hittable: side1, Material: material})

	side2 := geometries.NewQuad(
		core.NewVec3(472.0, 0.0, 406.0),
		core.NewVec3(472.0, 330.0, 406.0),
		core.NewVec3(314.0, 330.0, 456.0),
		core.NewVec3(314.0, 0.0, 456.0))
	objects = append(objects, scene.Object{Hittable: side2, Material: material})

	side3 := geometries.NewQuad(
		core.NewVec3(314.0, 0.0, 456.0),
		core.NewVec3(314.0, 330.0, 456.0),
		core.NewVec3(265.0, 330.0, 296.0),
		core.NewVec3(265.0, 0.0, 296.0))
	objects = append(objects, scene.Object{Hittable: side3, Material: material})

	side4 := geometries.NewQuad(
		core.NewVec3(265.0, 0.0, 296.0),
		core.NewVec3(265.0, 330.0, 296.0),
		core.NewVec3(423.0, 330.0, 247.0),
		core.NewVec3(423.0, 0.0, 247.0))
	objects = append(objects, scene.Object{Hittable: side4, Material: material})

	return objects
}
