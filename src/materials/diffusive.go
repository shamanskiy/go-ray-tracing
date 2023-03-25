package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
)

type Diffusive struct {
	color color.Color
}

func NewDiffusive(color color.Color) Diffusive {
	return Diffusive{color}
}

func (d Diffusive) Reflect(ray core.Ray, hit objects.HitRecord) *Reflection {
	reflectedDirection := hit.Normal.Add(core.Random().VecInUnitSphere())
	return &Reflection{core.NewRay(hit.Point, reflectedDirection), d.color}
}
