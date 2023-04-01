package objects

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/slices"
)

type Sphere struct {
	center core.Vec3
	radius core.Real
}

func NewSphere(center core.Vec3, radius core.Real) Sphere {
	return Sphere{
		center: center,
		radius: radius,
	}
}

func (s Sphere) TestRay(ray core.Ray) []core.Real {
	centerToOrigin := ray.Origin().Sub(s.center)

	a := ray.Direction().Dot(ray.Direction())
	b := 2.0 * ray.Direction().Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - s.radius*s.radius
	solution := core.SolveQuadEquation(a, b, c)

	if solution.NoSolution {
		return []core.Real{}
	}

	solutions := []core.Real{solution.Left, solution.Right}
	return slices.Filter(solutions, func(value core.Real) bool {
		return value >= 0
	})
}

func (s Sphere) EvaluateHit(ray core.Ray, hitParam core.Real) HitRecord {
	hitPoint := ray.Eval(hitParam)
	hitNormal := hitPoint.Sub(s.center).Div(s.radius)
	return HitRecord{hitPoint, hitNormal}
}

func (s Sphere) Center() core.Vec3 {
	return s.center
}

func (s Sphere) Radius() core.Real {
	return s.radius
}
