package objects

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
)

type Sphere struct {
	Center core.Vec3
	Radius core.Real
}

func (s Sphere) HitWithMin(ray core.Ray, minParam core.Real) core.HitRecord {
	centerToOrigin := ray.Origin.Sub(s.Center)

	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - s.Radius*s.Radius
	left, right, err := core.SolveQuadraticEquation(a, b, c)

	var hit core.HitRecord
	if err != nil {
		return hit
	}

	if left >= minParam {
		hit.Param = left
		hit.Hit = true
	}

	if !hit.Hit && right >= minParam {
		hit.Param = right
		hit.Hit = true
	}

	if hit.Hit {
		hit.Point = ray.Eval(hit.Param)
		hit.Normal = hit.Point.Sub(s.Center).Mul(1 / s.Radius)
	}

	return hit
}

func (s Sphere) Hit(ray core.Ray) core.HitRecord {
	return s.HitWithMin(ray, 0.0)
}
