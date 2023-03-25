package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
)

type Diffusive struct {
	color      color.Color
	randomizer random.RandomGenerator
}

func NewDiffusive(color color.Color, randomizer random.RandomGenerator) Diffusive {
	return Diffusive{color, randomizer}
}

func (d Diffusive) Color() color.Color {
	return d.color
}

func (d Diffusive) Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) *Reflection {
	reflectedDirection := normalAtHitPoint.Add(d.randomizer.Vec3InUnitSphere())
	return &Reflection{core.NewRay(hitPoint, reflectedDirection), d.color}
}
