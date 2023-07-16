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
	Point    core.Vec3
	Normal   core.Vec3
	Material materials.Material
}
