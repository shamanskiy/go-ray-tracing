package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/google/uuid"
)

type Plane struct {
	origin core.Vec3
	normal core.Vec3
	id     uuid.UUID
}

func NewPlane(origin, normal core.Vec3) Plane {
	return Plane{
		origin: origin,
		normal: normal,
		id:     uuid.New(),
	}
}

// R = A + Bt
// Plane: dot(X-Origin, N) = 0
// t = dot(Origin-A,N) / dot(B,N)
func (p Plane) TestRay(ray core.Ray, params core.Interval) Hit {
	dotBN := ray.Direction().Dot(p.normal)
	if core.Abs(dotBN) <= core.Tolerance {
		return Hit{}
	}

	rayPlaneDistance := p.origin.Sub(ray.Origin()).Dot(p.normal)
	hitParam := rayPlaneDistance / dotBN

	if params.Contains(hitParam) {
		return Hit{true, hitParam, p}
	}

	return Hit{}
}

func (p Plane) EvaluateHit(ray core.Ray, hitParam core.Real) HitPoint {
	positiveSide := ray.Direction().Dot(p.normal) < 0
	var normal core.Vec3
	if positiveSide {
		normal = p.normal
	} else {
		normal = p.normal.Mul(-1)
	}

	return HitPoint{
		Point:  ray.Eval(hitParam),
		Normal: normal,
	}
}

func (p Plane) BoundingBox() core.Box {
	return core.NewInfiniteBox()
}

func (p Plane) Id() uuid.UUID {
	return p.id
}
