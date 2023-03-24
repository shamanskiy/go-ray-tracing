package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
)

type Reflection struct {
	Ray         core.Ray
	Attenuation color.Color
}

type Material interface {
	Reflect(ray core.Ray, hit objects.HitRecord) *Reflection
}
