package background

import (
	"fmt"

	"github.com/chewxy/math32"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type VerticalGradient struct {
	topColor    color.Color
	bottomColor color.Color
}

func NewVerticalGradient(bottomColor, topColor color.Color) VerticalGradient {
	return VerticalGradient{
		topColor:    topColor,
		bottomColor: bottomColor,
	}
}

func (g VerticalGradient) ColorRay(ray core.Ray) color.Color {
	normalizedDirection := ray.Direction().Normalize()

	if math32.IsNaN(normalizedDirection.Y()) {
		panic(fmt.Errorf("got ray with zero length direction"))
	}

	t := 0.5 * (normalizedDirection.Y() + 1.0)

	return color.Interpolate(g.bottomColor, g.topColor, t)
}

func (g VerticalGradient) BottomColor() color.Color {
	return g.bottomColor
}

func (g VerticalGradient) TopColor() color.Color {
	return g.topColor
}
