package geometries

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
	"github.com/google/uuid"
)

type Triangle struct {
	a, b, c core.Vec3
	id      uuid.UUID
}

func NewTriangle(a, b, c core.Vec3) Triangle {
	return Triangle{a, b, c, uuid.New()}
}

func (t Triangle) BoundingBox() core.Box {
	return core.NewBox(core.Vec3Min(core.Vec3Min(t.a, t.b), t.c),
		core.Vec3Max(core.Vec3Max(t.a, t.b), t.c))
}

func (t Triangle) TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit] {
	return optional.Empty[Hit]()
}
