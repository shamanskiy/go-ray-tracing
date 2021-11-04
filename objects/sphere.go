package objects

import (
	"github.com/Shamanskiy/go-ray-tracer/core"
)

type Sphere struct {
	Center core.Vec3
	Radius core.Real
}

func (s Sphere) HitWithMin(ray core.Ray, minParam core.Real) *HitRecord {
	centerToOrigin := ray.Origin.Sub(s.Center)

	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - s.Radius*s.Radius
	left, right, err := core.SolveQuadraticEquation(a, b, c)

	if err != nil {
		return nil
	}

	if left >= minParam {
		hitPoint := ray.Eval(left)
		hitNormal := hitPoint.Sub(s.Center).Mul(1 / s.Radius)
		return &HitRecord{left, hitPoint, hitNormal}
	}

	if right >= minParam {
		hitPoint := ray.Eval(right)
		hitNormal := hitPoint.Sub(s.Center).Mul(1 / s.Radius)
		return &HitRecord{right, hitPoint, hitNormal}
	}

	return nil
}

func (s Sphere) Hit(ray core.Ray) *HitRecord {
	return s.HitWithMin(ray, 0.0)
}
