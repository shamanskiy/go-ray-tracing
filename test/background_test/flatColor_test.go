package background_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/background"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func TestFlatColor_ShouldReturnColor(t *testing.T) {
	flatColor := background.NewFlatColor(anyColor)
	ray := core.NewRay(anyPoint, anyDirection)

	rayColor := flatColor.ColorRay(ray)

	assert.Equal(t, flatColor.Color(), rayColor)
}
