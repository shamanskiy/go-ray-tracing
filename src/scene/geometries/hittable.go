package geometries

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type HitPoint struct {
	Point  core.Vec3
	Normal core.Vec3
}

type Hit struct {
	HasHit          bool
	Param core.Real
}

type Hittable interface {
	TestRay(ray core.Ray, params core.Interval) Hit
	EvaluateHit(ray core.Ray, hitParam core.Real) HitPoint
	BoundingBox() core.Box
}
