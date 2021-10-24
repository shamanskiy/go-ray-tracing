package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
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
			r := float64(x) / float64(width)
			g := float64(y) / float64(height)
			img.Set(x, y, color.RGBA{uint8(math.Round(255.99 * r)), uint8(math.Round(255.99 * g)), uint8(math.Round(255.99 * 0.2)), 0xff})
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
