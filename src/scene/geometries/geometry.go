package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/google/uuid"
)

type HitPoint struct {
	Point  core.Vec3
	Normal core.Vec3
}

type Geometry interface {
	Hittable
	EvaluateHit(ray core.Ray, hitParam core.Real) HitPoint
	Id() uuid.UUID
}
