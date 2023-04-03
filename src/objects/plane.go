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

func (p Plane) Normal() core.Vec3 {
	return p.normal
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

	return slices.Filter(t, slices.GreaterOrEqualThan(core.Real(0.)))
}

func (p Plane) EvaluateHit(ray core.Ray, hitParam core.Real) HitRecord {
	positiveSide := ray.Direction().Dot(p.normal) < 0
	var normal core.Vec3
	if positiveSide {
		normal = p.normal
	} else {
		normal = p.normal.Mul(-1)
	}

	return HitRecord{
		Point:  ray.Eval(hitParam),
		Normal: normal,
	}
}
