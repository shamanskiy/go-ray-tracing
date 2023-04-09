package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
)

type ReflectionType int

const (
	Scattered ReflectionType = iota
	Emitted
	Absorbed
)

type Reflection struct {
	Type  ReflectionType
	Ray   core.Ray
	Color color.Color
}

type Material interface {
	Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) Reflection
}
