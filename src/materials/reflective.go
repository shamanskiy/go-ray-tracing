package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
)

type Reflective struct {
	Color     color.Color
	Fuzziness core.Real
}

func NewReflectiveFuzzy(color color.Color, fuzziness core.Real) Reflective {
	if fuzziness < 0 {
		fuzziness = 0
	}
	if fuzziness > 1 {
		fuzziness = 1
	}
	return Reflective{color, fuzziness}
}

func (r Reflective) Reflect(ray core.Ray, hit objects.HitRecord) *Reflection {
	reflectedDirection := ray.Direction.Normalize().Reflect(hit.Normal)
	fuzzyPerturbation := core.Random().VecInUnitSphere().Mul(r.Fuzziness)
	reflectedDirection = reflectedDirection.Add(fuzzyPerturbation)

	if reflectedDirection.Dot(hit.Normal) > 0 {
		return &Reflection{core.Ray{hit.Point, reflectedDirection}, r.Color}
	} else {
		return nil
	}
}
