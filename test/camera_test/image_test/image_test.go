package image_test

import (
	rgba "image/color"
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/camera/image"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/stretchr/testify/assert"
)

const (
	IMAGE_WIDTH  = 2
	IMAGE_HEIGHT = 1
)

func TestImage_ShouldConvertToRGBA(t *testing.T) {
	image := image.NewImage(IMAGE_WIDTH, IMAGE_HEIGHT)
	image.SetPixelColor(0, 0, color.Red)
	image.SetPixelColor(1, 0, color.Blue)

	rgbaImage := image.ConvertToRGBA()

	assert.Equal(t, IMAGE_WIDTH, rgbaImage.Bounds().Max.X)
	assert.Equal(t, IMAGE_HEIGHT, rgbaImage.Bounds().Max.Y)
	assert.Equal(t, rgba.RGBA{255, 0, 0, 255}, rgba.Black, rgbaImage.At(0, 0))
	assert.Equal(t, rgba.RGBA{0, 0, 255, 255}, rgbaImage.At(1, 0))
}
