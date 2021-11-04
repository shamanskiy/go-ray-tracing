package core

import "github.com/Shamanskiy/go-ray-tracer/utils"

type Sphere struct {
	Center Vec3
	Radius Real
}

func (s Sphere) HitWithMin(ray Ray, minParam Real) HitRecord {
	centerToOrigin := ray.Origin.Sub(s.Center)

	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - s.Radius*s.Radius
	left, right, err := utils.SolveQuadraticEquation(a, b, c)

	var hit HitRecord
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

func (s Sphere) Hit(ray Ray) HitRecord {
	return s.HitWithMin(ray, 0.0)
}
