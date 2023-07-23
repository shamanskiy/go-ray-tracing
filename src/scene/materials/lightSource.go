package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type DiffusiveLight struct {
	color     color.Color
	intensity core.Real
}

func NewDiffusiveLight(color color.Color, intensity core.Real) DiffusiveLight {
	return DiffusiveLight{color: color, intensity: intensity}
}

func (d DiffusiveLight) Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) Reflection {
	return Reflection{
		Type:  Emitted,
		Color: d.color.Mul(d.intensity),
	}
}
