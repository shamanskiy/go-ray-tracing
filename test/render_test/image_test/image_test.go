package image_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/render/image"
	"github.com/stretchr/testify/assert"
)

const (
	IMAGE_WIDTH  = 5
	IMAGE_HEIGHT = 4
)

func TestImage_ShouldSetPixelColor(t *testing.T) {
	image := image.NewImage(IMAGE_WIDTH, IMAGE_HEIGHT)
	pixelColor := color.Red

	image.SetPixelColor(4, 3, pixelColor)

	assert.Equal(t, pixelColor, image.PixelColor(4, 3))
}
