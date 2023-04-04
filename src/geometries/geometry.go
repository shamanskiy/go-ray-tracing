package geometries

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type HitPoint struct {
	Point  core.Vec3
	Normal core.Vec3
}

type Geometry interface {
	TestRay(ray core.Ray) (hitParams []core.Real)
	EvaluateHit(ray core.Ray, hitParam core.Real) HitPoint
}
