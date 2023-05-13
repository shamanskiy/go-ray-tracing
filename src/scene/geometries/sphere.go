package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
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
		return optional.Of(Hit{solution.Left, sphere})
	}

	if params.Contains(solution.Right) {
		return optional.Of(Hit{solution.Right, sphere})
	}

	return optional.Empty[Hit]()
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
