package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

type Reflective struct {
	Color core.Color
}

func (r Reflective) Reflect(ray core.Ray, hit objects.HitRecord) *Reflection {
	reflectedDirection := core.Reflect(ray.Direction, hit.Normal)

	if reflectedDirection.Dot(hit.Normal) > 0 {
		return &Reflection{core.Ray{hit.Point, reflectedDirection}, r.Color}
	} else {
		return nil
	}
}
