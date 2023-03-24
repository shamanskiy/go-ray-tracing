package color_test

import (
	rgba "image/color"

	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/stretchr/testify/assert"
)

func TestColor_ToRGBA(t *testing.T) {
	t.Log("Black to RGB")
	colorIn := color.NewColor(0, 0, 0)
	colorOut := rgba.RGBA{0, 0, 0, 255}
	assert.Equal(t, colorOut, colorIn.ToRGBA())

	t.Log("White to RGB")
	colorIn = color.NewColor(1., 1., 1.)
	colorOut = rgba.RGBA{255, 255, 255, 255}
	assert.Equal(t, colorOut, colorIn.ToRGBA())

	t.Log("Gray to RGB with gamma correction")
	colorIn = color.NewColor(0.64, 0.64, 0.64)
	colorOut = rgba.RGBA{204, 204, 204, 255}
	assert.Equal(t, colorOut, colorIn.ToRGBA())
}
