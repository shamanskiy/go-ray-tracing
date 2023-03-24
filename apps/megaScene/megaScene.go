// This is a big scene with three big central spheres and
// 450+ small randomly generated spheres.
// ATTENTION: the scene can take a while to render
package main

import (
	"image"
	"image/png"
	"os"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/Shamanskiy/go-ray-tracer/src/render"
)

func tooCloseToBigSpheres(sphere objects.Sphere, bigSpheres []objects.Sphere) bool {
	for _, bigSphere := range bigSpheres {
		distance := sphere.Center.Sub(bigSphere.Center).Len()
		if distance < 1.05*sphere.Radius+bigSphere.Radius {
			return true
		}
	}
	return false
}

func genRandomSmallSphere(i, j int) objects.Sphere {
	center := core.Vec3{
		core.Real(i) + 0.8*core.Random().From01(),
		0.2,
		core.Real(j) + 0.8*core.Random().From01(),
	}
	return objects.Sphere{Center: center, Radius: 0.2}
}

func genRandomMaterial() materials.Material {
	materialChoice := core.Random().From01()

	if materialChoice < 0.7 {
		materialColor := core.Random().Vec3From01()
		return materials.Diffusive{materialColor}
	} else if materialChoice < 0.9 {
		metallicColor := core.Vec3{1., 1., 1.}.Add(core.Random().Vec3From01()).Mul(0.5)
		fuzziness := 0.1 * core.Random().From01()
		return materials.NewReflectiveFuzzy(metallicColor, fuzziness)
	} else {
		return materials.NewTransparent(1.5)
	}
}

func makeScene() *render.Scene {
	scene := render.Scene{SkyColorTop: core.SkyBlue, SkyColorBottom: core.White}

	// Huge sphere = floor
	scene.Add(objects.Sphere{Center: core.Vec3{0.0, -1000., 0.0}, Radius: 1000.0},
		materials.Diffusive{core.GrayMedium})

	// Three main spheres
	sphereRed := objects.Sphere{Center: core.Vec3{-2.5, 1.0, 1}, Radius: 1.0}
	scene.Add(sphereRed, materials.Diffusive{core.Red})
	sphereGlass := objects.Sphere{Center: core.Vec3{0.0, 1.0, 0.15}, Radius: 1.0}
	scene.Add(sphereGlass, materials.NewTransparent(1.5))
	sphereMirror := objects.Sphere{Center: core.Vec3{2.5, 1.0, 0.0}, Radius: 1.0}
	scene.Add(sphereMirror, materials.NewReflectiveFuzzy(core.GrayLight, 0.03))

	// Little spheres randomly offset from a regular grid
	gridSize := 11
	bigSpheres := []objects.Sphere{sphereRed, sphereGlass, sphereMirror}
	for i := -gridSize; i < gridSize; i++ {
		for j := -gridSize; j < gridSize; j++ {
			sphere := genRandomSmallSphere(i, j)
			if tooCloseToBigSpheres(sphere, bigSpheres) {
				continue
			}

			material := genRandomMaterial()
			scene.Add(sphere, material)
		}
	}

	return &scene
}

func makeCamera() *render.Camera {
	settings := render.DefaultCameraSettings()
	settings.AspectRatio = 16. / 9.
	settings.ImagePixelHeight = 300
	settings.Antialiasing = 10
	settings.LookFrom = core.Vec3{3.5, 1.35, 1.9}
	settings.LookAt = core.Vec3{3., 1.25, 1.5}
	settings.VerticalFOV = 75

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
	saveImage(img, "megaScene.png")
}
