package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/chewxy/math32"
	"github.com/schollz/progressbar/v3"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/materials"
	"github.com/Shamanskiy/go-ray-tracer/objects"
	"github.com/Shamanskiy/go-ray-tracer/render"
)

func main() {
	scene := render.Scene{SkyColorTop: core.SkyBlue, SkyColorBottom: core.White}
	scene.Add(objects.Sphere{Center: core.Vec3{0.0, 0.0, -1.0}, Radius: 0.5},
		materials.Diffusive{core.Red})
	scene.Add(objects.Sphere{Center: core.Vec3{-1.0, 0.0, -1.0}, Radius: 0.5},
		materials.Reflective{Color: core.GrayLight})
	scene.Add(objects.Sphere{Center: core.Vec3{1.0, 0.0, -1.0}, Radius: 0.5},
		materials.NewReflectiveWithFuzziness(core.Golden, 1.0))
	scene.Add(objects.Sphere{Center: core.Vec3{0.0, -100.5, -1.0}, Radius: 100.0},
		materials.Diffusive{core.GrayMedium})

	width := 800
	height := 400
	sampling := 1

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	camera := render.Camera{
		Origin:            core.Vec3{0.0, 0.0, 0.0},
		Upper_left_corner: core.Vec3{-2.0, 1.0, -1.0},
		Horizontal:        core.Vec3{4.0, 0.0, 0.0},
		Vertical:          core.Vec3{0.0, -2.0, 0.0},
	}

	bar := progressbar.Default(int64(width), "rendering")

	start := time.Now()
	for x := 0; x < width; x++ {
		bar.Add(1)
		for y := 0; y < height; y++ {
			var clr core.Color
			for s := 0; s < sampling; s++ {
				u := (core.Real(x) + rand.Float32()) / core.Real(width)
				v := (core.Real(y) + rand.Float32()) / core.Real(height)
				ray := camera.GetRay(u, v)
				c := scene.TestRay(ray)
				clr = clr.Add(c)
			}
			clr = clr.Mul(1.0 / core.Real(sampling))
			img.Set(x, y, toRGBA(clr))
		}
	}
	elapsed := time.Since(start)
	log.Printf("Rendering took %s", elapsed)

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func toRGBA(c core.Color) color.RGBA {
	return color.RGBA{toZero255(c.X()), toZero255(c.Y()), toZero255(c.Z()), 0xff}
}

func toZero255(x core.Real) uint8 {
	return uint8(math32.Floor(255.99 * math32.Sqrt(x)))
}
