package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/chewxy/math32"

	"github.com/Shamanskiy/go-ray-tracer/core"
)

func main() {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	origin := core.Vec3{0.0, 0.0, 0.0}
	upper_left_corner := core.Vec3{-2.0, 1.0, -1.0}
	horizontal := core.Vec3{4.0, 0.0, 0.0}
	vertical := core.Vec3{0.0, -2.0, 0.0}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			u := float32(x) / float32(width)
			v := float32(y) / float32(height)
			ray := core.Ray{origin, upper_left_corner.Add(horizontal.Mul(u)).Add(vertical.Mul(v))}

			c := testRay(ray)
			img.Set(x, y, toRGBA(c))
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func toRGBA(c core.Color) color.RGBA {
	return color.RGBA{toZero255(c.X()), toZero255(c.Y()), toZero255(c.Z()), 0xff}
}

func toZero255(x float32) uint8 {
	return uint8(math32.Floor(255.99 * x))
}

func testRay(ray core.Ray) core.Color {
	if hitSphere(core.Vec3{0.0, 0.0, -1.0}, 0.5, ray) {
		return core.Color{1.0, 0.0, 0.0}
	}

	unit_direction := ray.Direction.Normalize()
	t := 0.5 * (unit_direction.Y() + 1.0)
	A := core.Color{1.0, 1.0, 1.0}.Mul(1.0 - t)
	B := core.Color{0.5, 0.7, 1.0}.Mul(t)
	return A.Add(B)
}

func hitSphere(center core.Vec3, radius core.Real, ray core.Ray) bool {
	oc := ray.Origin.Sub(center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(oc)
	c := oc.Dot(oc) - radius*radius
	d := b*b - 4*a*c
	return d > 0.0
}
