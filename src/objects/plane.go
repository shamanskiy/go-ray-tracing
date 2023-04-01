package objects

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/slices"
)

type Plane struct {
	origin core.Vec3
	normal core.Vec3
}

func NewPlane(origin, normal core.Vec3) Plane {
	return Plane{
		origin: origin,
		normal: normal,
	}
}

// R = A + Bt
// Plane: dot(X-Origin, N) = 0
// t = dot(Origin-A,N) / dot(B,N)
func (p Plane) TestRay(ray core.Ray) (hitParams []core.Real) {
	dotBN := ray.Direction().Dot(p.normal)
	if core.Abs(dotBN) <= core.Tolerance {
		return []core.Real{}
	}

	rayPlaneDistance := p.origin.Sub(ray.Origin()).Dot(p.normal)
	t := []core.Real{rayPlaneDistance / dotBN}

	return slices.Filter(t, func(value core.Real) bool {
		return value >= 0
	})
}

func (p Plane) EvaluateHit(ray core.Ray, hitParam core.Real) HitRecord {
	return HitRecord{
		Point:  ray.Eval(hitParam),
		Normal: p.normal,
	}
}
