package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
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

func (sphere Sphere) TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit] {
	centerToOrigin := ray.Origin().Sub(sphere.center)

	a := ray.Direction().Dot(ray.Direction())
	b := 2.0 * ray.Direction().Dot(centerToOrigin)
	c := centerToOrigin.Dot(centerToOrigin) - sphere.radius*sphere.radius
	solution := core.SolveQuadEquation(a, b, c)

	if solution.NoSolution {
		return optional.Empty[Hit]()
	}

	if params.Contains(solution.Left) {
		return optional.Of(sphere.evaluateHit(ray, solution.Left))
	}

	if params.Contains(solution.Right) {
		return optional.Of(sphere.evaluateHit(ray, solution.Right))
	}

	return optional.Empty[Hit]()
}

func (sphere Sphere) evaluateHit(ray core.Ray, hitParam core.Real) Hit {
	hitPoint := ray.Eval(hitParam)
	return Hit{
		Param:  hitParam,
		Point:  hitPoint,
		Normal: hitPoint.Sub(sphere.center).Div(sphere.radius),
	}
}

func (sphere Sphere) InContactWith(other Sphere) bool {
	distance := sphere.center.Sub(other.center).Len()
	return distance <= sphere.radius+other.radius
}

func (sphere Sphere) BoundingBox() core.Box {
	centerToCorner := core.NewVec3(sphere.radius, sphere.radius, sphere.radius)
	return core.NewBox(sphere.center.Sub(centerToCorner), sphere.center.Add(centerToCorner))
}
