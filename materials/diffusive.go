package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

type Diffusive struct {
	Color core.Color
}

func (d Diffusive) Reflect(ray core.Ray, hit objects.HitRecord) *Reflection {
	reflectedDirection := hit.Normal.Add(core.Random().VecInUnitSphere())
	return &Reflection{core.Ray{hit.Point, reflectedDirection}, d.Color}
}
