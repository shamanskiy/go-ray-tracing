package background_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/stretchr/testify/assert"
)

var RAY_ORIGIN = core.NewVec3(999, 666, 333)
var RAY_DIRECTION = core.NewVec3(1, 2, 3)
var BACKGROUND_COLOR = color.White
var OTHER_BACKGROUND_COLOR = color.Black

func TestVerticalGradient_ShouldBlendColorsForHorizontalRay(t *testing.T) {
	gradient := background.NewVerticalGradient(BACKGROUND_COLOR, OTHER_BACKGROUND_COLOR)
	ray := core.NewRay(RAY_ORIGIN, core.NewVec3(0, 0, 1))

	rayColor := gradient.ColorRay(ray)

	expectedColor := color.Interpolate(BACKGROUND_COLOR, OTHER_BACKGROUND_COLOR, 0.5)
	assert.Equal(t, expectedColor, rayColor)
}

func TestVerticalGradient_ShouldReturnBottomColorForVerticalDownRay(t *testing.T) {
	gradient := background.NewVerticalGradient(BACKGROUND_COLOR, OTHER_BACKGROUND_COLOR)
	ray := core.NewRay(RAY_ORIGIN, core.NewVec3(0, -1, 0))

	rayColor := gradient.ColorRay(ray)

	assert.Equal(t, BACKGROUND_COLOR, rayColor)
}

func TestVerticalGradient_ShouldReturnTopColorForVerticalUpRay(t *testing.T) {
	gradient := background.NewVerticalGradient(BACKGROUND_COLOR, OTHER_BACKGROUND_COLOR)
	ray := core.NewRay(RAY_ORIGIN, core.NewVec3(0, 1, 0))

	rayColor := gradient.ColorRay(ray)

	assert.Equal(t, OTHER_BACKGROUND_COLOR, rayColor)
}
