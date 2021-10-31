package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/chewxy/math32"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := mgl32.Vec3{float32(x) / float32(width), float32(y) / float32(height), 0.5}
			img.Set(x, y, color.RGBA{toZero255(c.X()), toZero255(c.Y()), toZero255(c.Z()), 0xff})
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func toZero255(x float32) uint8 {
	return uint8(math32.Floor(255.99 * x))
}
