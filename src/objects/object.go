package objects

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type HitRecord struct {
	Point  core.Vec3
	Normal core.Vec3
}

type Object interface {
	TestRay(ray core.Ray) (hitParams []core.Real)
	EvaluateHit(ray core.Ray, hitParam core.Real) HitRecord
}
