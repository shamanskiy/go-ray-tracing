package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

type Reflection struct {
	Ray         core.Ray
	Attenuation core.Color
}

type Material interface {
	Reflect(ray core.Ray, hit objects.HitRecord) *Reflection
}
