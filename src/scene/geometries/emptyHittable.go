package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
)

type emptyHittable struct{}

func (e emptyHittable) TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit] {
	return optional.Empty[Hit]()
}

func (e emptyHittable) BoundingBox() core.Box {
	return core.NewEmptyBox()
}
