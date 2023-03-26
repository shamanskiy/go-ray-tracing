package objects

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
)

type Sphere struct {
	Center core.Vec3
	Radius core.Real
}

func (s Sphere) TestRay(ray core.Ray) []core.Real {
	centerToOrigin := ray.Origin().Sub(s.Center)

	a := ray.Direction().Dot(ray.Direction())
	b := 2.0 * ray.Direction().Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - s.Radius*s.Radius
	solution := core.SolveQuadEquation(a, b, c)

	if solution.NoSolution {
		return []core.Real{}
	}

	return []core.Real{solution.Left, solution.Right}
}

func (s Sphere) EvaluateHit(ray core.Ray, hitParam core.Real) HitRecord {
	hitPoint := ray.Eval(hitParam)
	hitNormal := hitPoint.Sub(s.Center).Div(s.Radius)
	return HitRecord{hitPoint, hitNormal}
}
