package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
)

type Diffusive struct {
	color      color.Color
	randomizer random.RandomGenerator
}

func NewDiffusive(color color.Color, randomizer random.RandomGenerator) Diffusive {
	return Diffusive{color, randomizer}
}

func (d Diffusive) Reflect(ray core.Ray, hit objects.HitRecord) *Reflection {
	reflectedDirection := hit.Normal.Add(d.randomizer.Vec3InUnitSphere())
	return &Reflection{core.NewRay(hit.Point, reflectedDirection), d.color}
}
