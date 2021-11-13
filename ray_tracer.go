package main

import (
	"fmt"
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
	scene := render.Scene{}
	scene.Add(objects.Sphere{Center: core.Vec3{0.0, 0.0, -1.0}, Radius: 0.5})
	scene.Add(objects.Sphere{Center: core.Vec3{0.0, -100.5, -1.0}, Radius: 100.0})

	width := 800
	height := 400
	sampling := 10

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	camera := render.Camera{
		Origin:            core.Vec3{0.0, 0.0, 0.0},
		Upper_left_corner: core.Vec3{-2.0, 1.0, -1.0},
		Horizontal:        core.Vec3{4.0, 0.0, 0.0},
		Vertical:          core.Vec3{0.0, -2.0, 0.0},
	}

	bar := progressbar.NewOptions(width,
		progressbar.OptionFullWidth(),
		progressbar.OptionSetDescription("rendering"),
		progressbar.OptionShowCount(),
		progressbar.OptionOnCompletion(func() { fmt.Println() }),
		progressbar.OptionShowIts(),
		progressbar.OptionSetItsString("col"),
		progressbar.OptionThrottle(100*time.Millisecond))

	start := time.Now()
	for x := 0; x < width; x++ {
		bar.Add(1)
		for y := 0; y < height; y++ {
			var clr core.Color
			for s := 0; s < sampling; s++ {
				u := (float32(x) + rand.Float32()) / float32(width)
				v := (float32(y) + rand.Float32()) / float32(height)
				ray := camera.GetRay(u, v)
				c := testRay(ray, &scene, 0)
				clr = clr.Add(c)
			}
			clr = clr.Mul(1.0 / float32(sampling))
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

func toZero255(x float32) uint8 {
	return uint8(math32.Floor(255.99 * math32.Sqrt(x)))
}

func testRay(ray core.Ray, scene *render.Scene, depth int) core.Color {
	hit := scene.HitWithMin(ray, 0.0001)
	m := materials.Diffusive{core.Color{0.5, 0.5, 0.5}}
	if hit != nil {
		reflection := m.Reflect(ray, *hit)
		if reflection != nil && depth < 10 {
			return core.MulElem(testRay(reflection.Ray, scene, depth+1), reflection.Attenuation)
		} else {
			return core.Color{0.0, 0.0, 0.0}
		}
	}

	unit_direction := ray.Direction.Normalize()
	t := 0.5 * (unit_direction.Y() + 1.0)
	A := core.White.Mul(1.0 - t)
	B := core.SkyBlue.Mul(t)
	return A.Add(B)
}
