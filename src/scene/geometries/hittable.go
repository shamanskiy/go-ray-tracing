package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
	"github.com/google/uuid"
)

type HitPoint struct {
	Point  core.Vec3
	Normal core.Vec3
}

type Hit struct {
	Param       core.Real
	HitGeometry Geometry
}

type Hittable interface {
	TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit]
	BoundingBox() core.Box
}

type Geometry interface {
	Hittable
	EvaluateHit(ray core.Ray, hitParam core.Real) HitPoint
	Id() uuid.UUID
}
