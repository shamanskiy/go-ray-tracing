package background_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func TestFlatColor_ShouldReturnColor(t *testing.T) {
	flatColor := background.NewFlatColor(BACKGROUND_COLOR)
	ray := core.NewRay(RAY_ORIGIN, RAY_DIRECTION)

	rayColor := flatColor.ColorRay(ray)

	assert.Equal(t, BACKGROUND_COLOR, rayColor)
}
