package materials

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/objects"
)

type Reflection struct {
	NotAbsorbed bool
	Ray         core.Ray
	Attenuation core.Real
}

type Material interface {
	Reflect(ray core.Ray, hit objects.HitRecord) Reflection
}
