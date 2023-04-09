package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type DiffusiveLight struct {
	color color.Color
}

func NewDiffusiveLight(color color.Color) DiffusiveLight {
	return DiffusiveLight{color}
}

func (d DiffusiveLight) Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) Reflection {
	return Reflection{
		Type:  Emitted,
		Color: d.color,
	}
}
