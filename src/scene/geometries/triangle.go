package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
)

type Triangle struct {
	v0, v1, v2 core.Vec3
	n0, n1, n2 core.Vec3
}

func NewTriangle(v0, v1, v2, n0, n1, n2 core.Vec3) Triangle {
	return Triangle{v0, v1, v2, n0, n1, n2}
}

func (t Triangle) BoundingBox() core.Box {
	return core.NewBox(core.Vec3Min(core.Vec3Min(t.v0, t.v1), t.v2),
		core.Vec3Max(core.Vec3Max(t.v0, t.v1), t.v2))
}

// https://en.wikipedia.org/wiki/M%C3%B6ller%E2%80%93Trumbore_intersection_algorithm
func (t Triangle) TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit] {
	edge1 := t.v1.Sub(t.v0)
	edge2 := t.v2.Sub(t.v0)
	rayCrossEdge2 := ray.Direction().Cross(edge2)
	det := edge1.Dot(rayCrossEdge2)

	if rayIsParallelToTriangle(det) {
		return optional.Empty[Hit]()
	}

	invDet := 1.0 / det
	originToVertex0 := ray.Origin().Sub(t.v0)

	u := invDet * originToVertex0.Dot(rayCrossEdge2)
	if !core.NewInterval(0, 1).Contains(u) {
		return optional.Empty[Hit]()
	}

	q := originToVertex0.Cross(edge1)
	v := invDet * ray.Direction().Dot(q)
	if !core.NewInterval(0, 1-u).Contains(v) {
		return optional.Empty[Hit]()
	}

	rayParam := invDet * edge2.Dot(q)
	if !params.Contains(rayParam) {
		return optional.Empty[Hit]()
	}

	return optional.Of(t.evaluateHit(ray, rayParam, u, v))
}

func rayIsParallelToTriangle(det core.Real) bool {
	return core.Abs(det) < core.Tolerance
}

func (t Triangle) evaluateHit(ray core.Ray, hitParam, u, v core.Real) Hit {
	hitPoint := ray.Eval(hitParam)

	return Hit{
		Param:  hitParam,
		Point:  hitPoint,
		Normal: t.normalGouraud(u, v),
	}
}

func (t Triangle) normalGouraud(u, v core.Real) core.Vec3 {
	return t.n0.Mul(1.0 - u - v).Add(t.n1.Mul(u)).Add(t.n2.Mul(v)).Normalize()
}
