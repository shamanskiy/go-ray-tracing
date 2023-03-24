package color_test

import (
	rgba "image/color"

	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/stretchr/testify/assert"
)

func TestColor_BlackToRGBA(t *testing.T) {
	assert.Equal(t, rgba.RGBA{0, 0, 0, 255}, color.Black.ToRGBA())
}

func TestColor_WhiteToRGBA(t *testing.T) {
	assert.Equal(t, rgba.RGBA{255, 255, 255, 255}, color.White.ToRGBA())
}

func TestColor_GrayToRGBA(t *testing.T) {
	assert.Equal(t, rgba.RGBA{181, 181, 181, 255}, color.GrayMedium.ToRGBA())
}
