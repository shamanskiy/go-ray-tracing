package objects

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
)

type Sphere struct {
	Center core.Vec3
	Radius core.Real
}

func (s Sphere) HitWithMin(ray core.Ray, minParam core.Real) *HitRecord {
	centerToOrigin := ray.Origin().Sub(s.Center)

	a := ray.Direction().Dot(ray.Direction())
	b := 2.0 * ray.Direction().Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - s.Radius*s.Radius
	solution := core.SolveQuadEquation(a, b, c)

	if solution.NoSolution {
		return nil
	}

	if solution.Left >= minParam {
		hitPoint := ray.Eval(solution.Left)
		hitNormal := hitPoint.Sub(s.Center).Mul(1 / s.Radius)
		return &HitRecord{solution.Left, hitPoint, hitNormal}
	}

	if solution.Right >= minParam {
		hitPoint := ray.Eval(solution.Right)
		hitNormal := hitPoint.Sub(s.Center).Mul(1 / s.Radius)
		return &HitRecord{solution.Right, hitPoint, hitNormal}
	}

	return nil
}

func (s Sphere) Hit(ray core.Ray) *HitRecord {
	return s.HitWithMin(ray, 0.0)
}
