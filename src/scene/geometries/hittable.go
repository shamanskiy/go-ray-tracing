package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
)

type Hittable interface {
	TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit]
	BoundingBox() core.Box
}

type Hit struct {
	Param    core.Real
	Geometry Geometry
	Material materials.Material
}

type HitPoint struct {
	Point  core.Vec3
	Normal core.Vec3
}

type Geometry interface {
	EvaluateHit(ray core.Ray, hitParam core.Real) HitPoint
}
