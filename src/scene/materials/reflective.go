package materials

import (
	"fmt"

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

func NewReflectiveFuzzy(color color.Color, fuzziness core.Real, randomizer random.RandomGenerator) Reflective {
	if fuzziness < 0 || fuzziness > 1 {
		panic(fmt.Errorf("fuzziness must be in range [0, 1], got %f", fuzziness))
	}
	return Reflective{
		color:      color,
		fuzziness:  fuzziness,
		randomizer: randomizer,
	}
}

func (r Reflective) Reflect(incidentDirection, hitPoint, normalAtHitPoint core.Vec3) Reflection {
	reflectedDirection := incidentDirection.Normalize().Reflect(normalAtHitPoint)
	fuzzyPerturbation := r.randomizer.Vec3InUnitSphere().Mul(r.fuzziness)
	reflectedDirection = reflectedDirection.Add(fuzzyPerturbation)

	if reflectedDirection.Dot(normalAtHitPoint) > 0 {
		return Reflection{
			Type:  Scattered,
			Ray:   core.NewRay(hitPoint, reflectedDirection),
			Color: r.color,
		}
	} else {
		return Reflection{Type: Absorbed}
	}
}
