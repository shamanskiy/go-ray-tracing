package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/google/uuid"
)

type Sphere struct {
	center core.Vec3
	radius core.Real
	id     uuid.UUID
}

func NewSphere(center core.Vec3, radius core.Real) Sphere {
	return Sphere{
		center: center,
		radius: radius,
		id:     uuid.New(),
	}
}

func (s Sphere) TestRay(ray core.Ray, params core.Interval) Hit {
	centerToOrigin := ray.Origin().Sub(s.center)

	a := ray.Direction().Dot(ray.Direction())
	b := 2.0 * ray.Direction().Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - s.radius*s.radius
	solution := core.SolveQuadEquation(a, b, c)

	if solution.NoSolution {
		return Hit{}
	}

	if params.Contains(solution.Left) {
		return Hit{true, solution.Left, s}
	}

	if params.Contains(solution.Right) {
		return Hit{true, solution.Right, s}
	}

	return Hit{}
}

func (s Sphere) EvaluateHit(ray core.Ray, hitParam core.Real) HitPoint {
	hitPoint := ray.Eval(hitParam)
	hitNormal := hitPoint.Sub(s.center).Div(s.radius)
	return HitPoint{hitPoint, hitNormal}
}

func (s Sphere) InContactWith(other Sphere) bool {
	distance := s.center.Sub(other.center).Len()
	return distance <= s.radius+other.radius
}

func (s Sphere) BoundingBox() core.Box {
	centerToCorner := core.NewVec3(s.radius, s.radius, s.radius)
	return core.NewBox(s.center.Sub(centerToCorner), s.center.Add(centerToCorner))
}

func (s Sphere) Id() uuid.UUID {
	return s.id
}
