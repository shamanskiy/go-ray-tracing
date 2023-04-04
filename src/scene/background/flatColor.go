package background

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type FlatColor struct {
	color color.Color
}

func NewFlatColor(color color.Color) FlatColor {
	return FlatColor{
		color: color,
	}
}

func (c FlatColor) ColorRay(ray core.Ray) color.Color {
	return c.color
}
