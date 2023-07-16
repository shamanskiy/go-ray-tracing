package geometries_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/geometries"
	"github.com/stretchr/testify/assert"
)

func TestBVHNode(t *testing.T) {
	sphereX := geometries.NewSphere(core.NewVec3(3, 0, 0), 1)
	sphereY := geometries.NewSphere(core.NewVec3(0, 3, 0), 1)
	sphereZ := geometries.NewSphere(core.NewVec3(0, 0, 3), 1)
	sphereCenter := geometries.NewSphere(core.NewVec3(0, 0, 0), 1)
	objects := []geometries.Hittable{sphereX, sphereY, sphereZ, sphereCenter}
	bvh := geometries.BuildBVH(objects)

	expectedBBox := core.NewBox(core.NewVec3(-1, -1, -1), core.NewVec3(4, 4, 4))
	assert.Equal(t, expectedBBox, bvh.BoundingBox())

	ray := core.NewRay(core.NewVec3(0, 0, 0), core.NewVec3(1, 0, 0))
	hit := bvh.TestRay(ray, core.NewInterval(2, 10))

	assert.EqualValues(t, 2, hit.Value().Param)
	assert.Equal(t, sphereX, hit.Value().Geometry)
}
