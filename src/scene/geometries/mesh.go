package geometries

import (
	"fmt"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
)

type Mesh struct {
	trianglesBHV *BVHNode
}

func NewMesh(triangles []Triangle) Mesh {
	hittables := make([]Hittable, 0, len(triangles))
	for _, triangle := range triangles {
		hittables = append(hittables, triangle)
	}

	return Mesh{
		trianglesBHV: BuildBVH(hittables),
	}
}

func (m Mesh) TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit] {
	return m.trianglesBHV.TestRay(ray, params)
}

func (m Mesh) BoundingBox() core.Box {
	return m.trianglesBHV.BoundingBox()
}

func NewQuad(a, b, c, d core.Vec3) Mesh {
	norm1 := core.Normal(a, b, c)
	norm2 := core.Normal(c, d, a)
	if !norm1.InDelta(norm2, core.Tolerance) {
		panic(fmt.Errorf("make quad: normals don't match: %v and %v", norm1, norm2))
	}
	triangle1 := NewTriangle(a, b, c)
	triangle2 := NewTriangle(c, d, a)

	return NewMesh([]Triangle{triangle1, triangle2})
}
