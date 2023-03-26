package materials

import (
	"fmt"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
)

type Transparent struct {
	color      color.Color
	refractor  RefractionCalculator
	randomizer random.RandomGenerator
}

func NewTransparent(refractionIndex core.Real, color color.Color, randomizer random.RandomGenerator) Transparent {
	if refractionIndex < 1 {
		panic(fmt.Errorf("refractionIndex must be at least 1, got %f", refractionIndex))
	}
	return Transparent{
		color:      color,
		refractor:  NewRefractionCalculator(refractionIndex),
		randomizer: randomizer}
}

func (m Transparent) Color() color.Color {
	return m.color
}

func (m Transparent) Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) *Reflection {
	refraction := m.refractor.Refract(incidentDirection, normalAtHitPoint)
	reflectedDirection := incidentDirection.Reflect(normalAtHitPoint)

	if refraction.FullInternalReflection() {
		return m.buildReflection(hitPoint, reflectedDirection)
	}

	// Transparent materials reflect a portion of the incoming light
	if m.randomizer.Real() > refraction.ReflectionRatio {
		return m.buildReflection(hitPoint, *refraction.Direction)
	} else {
		return m.buildReflection(hitPoint, reflectedDirection)
	}
}

func (m Transparent) buildReflection(hitPoint core.Vec3, direction core.Vec3) *Reflection {
	reflectedRay := core.NewRay(hitPoint, direction)
	return &Reflection{reflectedRay, m.color}
}
