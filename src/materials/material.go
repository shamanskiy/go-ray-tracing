package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type Reflection struct {
	Ray         core.Ray
	Attenuation color.Color
}

type Material interface {
	Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) *Reflection
}
