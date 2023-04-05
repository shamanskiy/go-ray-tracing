package image

import (
	"fmt"
	"image"

	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type Image struct {
	pixels [][]color.Color
}

func NewImage(width, height int) *Image {
	if width <= 0 || height <= 0 {
		panic(fmt.Errorf("new image: invalid size:  width %d, height %d", width, height))
	}

	pixels := make([][]color.Color, width)
	for column := range pixels {
		pixels[column] = make([]color.Color, height)
	}
	return &Image{pixels: pixels}
}

func (i *Image) Width() int {
	return len(i.pixels)
}

func (i *Image) Height() int {
	return len(i.pixels[0])
}

func (i *Image) SetPixelColor(x, y int, color color.Color) {
	i.pixels[x][y] = color
}

func (i *Image) ConvertToRGBA() *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{i.Width(), i.Height()}
	rgbaImage := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x, column := range i.pixels {
		for y, color := range column {
			rgbaImage.Set(x, y, color.ToRGBA())
		}
	}

	return rgbaImage
}
