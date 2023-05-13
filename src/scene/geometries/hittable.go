package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
)

type Hit struct {
	Param       core.Real
	HitGeometry Geometry
}

type Hittable interface {
	TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit]
	BoundingBox() core.Box
}

