package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

type Reflective struct {
	Color     core.Color
	fuzziness core.Real
}

func NewReflectiveFuzzy(color core.Color, fuzziness core.Real) Reflective {
	if fuzziness < 0 {
		fuzziness = 0
	}
	if fuzziness > 1 {
		fuzziness = 1
	}
	return Reflective{color, fuzziness}
}

func (r Reflective) Reflect(ray core.Ray, hit objects.HitRecord) *Reflection {
	reflectedDirection := core.Reflect(ray.Direction.Normalize(), hit.Normal)
	fuzzyPerturbation := core.Random().VecInUnitSphere().Mul(r.fuzziness)
	reflectedDirection = reflectedDirection.Add(fuzzyPerturbation)

	if reflectedDirection.Dot(hit.Normal) > 0 {
		return &Reflection{core.Ray{hit.Point, reflectedDirection}, r.Color}
	} else {
		return nil
	}
}
