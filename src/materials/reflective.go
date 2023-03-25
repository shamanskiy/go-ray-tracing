package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
)

type Reflective struct {
	color      color.Color
	fuzziness  core.Real
	randomizer random.RandomGenerator
}

func NewReflective(color color.Color, randomizer random.RandomGenerator) Reflective {
	return Reflective{
		color:      color,
		fuzziness:  0,
		randomizer: randomizer,
	}
}

func (r Reflective) Color() color.Color {
	return r.color
}

func (r Reflective) Fuzziness() core.Real {
	return r.fuzziness
}

func NewReflectiveFuzzy(color color.Color, fuzziness core.Real, randomizer random.RandomGenerator) Reflective {
	if fuzziness < 0 {
		fuzziness = 0
	}
	if fuzziness > 1 {
		fuzziness = 1
	}
	return Reflective{
		color:      color,
		fuzziness:  fuzziness,
		randomizer: randomizer,
	}
}

func (r Reflective) Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) *Reflection {
	reflectedDirection := incidentDirection.Normalize().Reflect(normalAtHitPoint)
	fuzzyPerturbation := r.randomizer.Vec3InUnitSphere().Mul(r.fuzziness)
	reflectedDirection = reflectedDirection.Add(fuzzyPerturbation)

	if reflectedDirection.Dot(normalAtHitPoint) > 0 {
		return &Reflection{core.NewRay(hitPoint, reflectedDirection), r.color}
	} else {
		return nil
	}
}
