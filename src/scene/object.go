package scene

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
)

type Object struct {
	Hittable geometries.Hittable
	Material materials.Material
}

func (o Object) TestRay(ray core.Ray, params core.Interval) optional.Optional[geometries.Hit] {
	optionalHit := o.Hittable.TestRay(ray, params)
	if optionalHit.Empty() {
		return optionalHit
	}

	hit := optionalHit.Value()
	hit.Material = o.Material
	return optional.Of(hit)
}

func (o Object) BoundingBox() core.Box {
	return o.Hittable.BoundingBox()
}
