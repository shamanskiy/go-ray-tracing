package background_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/stretchr/testify/assert"
)

var anyPoint = core.NewVec3(999, 666, 333)
var gradient = background.VerticalGradient{
	TopColor:    color.White,
	BottomColor: color.Black,
}

func TestVerticalGradient_ShouldBlendColorsForHorizontalRay(t *testing.T) {
	ray := core.NewRay(anyPoint, core.NewVec3(0, 0, 1))

	rayColor := gradient.ColorRay(ray)

	assert.Equal(t, color.GrayMedium, rayColor)
}

func TestVerticalGradient_ShouldReturnTopColorForVerticalUpRay(t *testing.T) {
	ray := core.NewRay(anyPoint, core.NewVec3(0, 1, 0))

	rayColor := gradient.ColorRay(ray)

	assert.Equal(t, gradient.TopColor, rayColor)
}

func TestVerticalGradient_ShouldReturnBottomColorForVerticalDownRay(t *testing.T) {
	ray := core.NewRay(anyPoint, core.NewVec3(0, -1, 0))

	rayColor := gradient.ColorRay(ray)

	assert.Equal(t, gradient.BottomColor, rayColor)
}
