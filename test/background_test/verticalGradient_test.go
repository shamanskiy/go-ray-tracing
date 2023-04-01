package background_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/stretchr/testify/assert"
)

var anyPoint = core.NewVec3(999, 666, 333)
var anyDirection = core.NewVec3(1, 2, 3)
var anyColor = color.White
var anyColor2 = color.Black
var equalColorMix = color.GrayMedium

func TestVerticalGradient_ShouldBlendColorsForHorizontalRay(t *testing.T) {
	gradient := background.NewVerticalGradient(anyColor, anyColor2)
	ray := core.NewRay(anyPoint, core.NewVec3(0, 0, 1))

	rayColor := gradient.ColorRay(ray)

	assert.Equal(t, equalColorMix, rayColor)
}

func TestVerticalGradient_ShouldReturnBottomColorForVerticalDownRay(t *testing.T) {
	gradient := background.NewVerticalGradient(anyColor, anyColor2)
	ray := core.NewRay(anyPoint, core.NewVec3(0, -1, 0))

	rayColor := gradient.ColorRay(ray)

	assert.Equal(t, gradient.BottomColor(), rayColor)
}

func TestVerticalGradient_ShouldReturnTopColorForVerticalUpRay(t *testing.T) {
	gradient := background.NewVerticalGradient(anyColor, anyColor2)
	ray := core.NewRay(anyPoint, core.NewVec3(0, 1, 0))

	rayColor := gradient.ColorRay(ray)

	assert.Equal(t, gradient.TopColor(), rayColor)
}
