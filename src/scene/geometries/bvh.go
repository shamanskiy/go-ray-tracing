package geometries

import (
	"math/rand"
	"sort"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/optional"
)

type BVHNode struct {
	boundingBox core.Box
	leftChild   Hittable
	rightChild  Hittable
}

func BuildBVH(geometries []Geometry) *BVHNode {
	if len(geometries) == 0 {
		panic("can't build bvh with zero geometries")
	}

	sortingAxis := rand.Intn(3)

	sort.Slice(geometries, func(i, j int) bool {
		return compareGeometries(geometries, i, j, sortingAxis)
	})

	bvhNode := &BVHNode{}

	if len(geometries) == 1 {
		bvhNode.leftChild = geometries[0]
		bvhNode.rightChild = geometries[0]
	} else if len(geometries) == 2 {
		bvhNode.leftChild = geometries[0]
		bvhNode.rightChild = geometries[1]
	} else {
		bvhNode.leftChild = BuildBVH(geometries[0 : len(geometries)/2])
		bvhNode.rightChild = BuildBVH(geometries[len(geometries)/2:])
	}

	bvhNode.boundingBox = bvhNode.leftChild.BoundingBox().Union(bvhNode.rightChild.BoundingBox())

	return bvhNode
}

func compareGeometries(geometries []Geometry, i, j, sortingAxis int) bool {
	minI := geometries[i].BoundingBox().Min().At(sortingAxis)
	minj := geometries[j].BoundingBox().Min().At(sortingAxis)
	return minI < minj
}

func (bhv *BVHNode) TestRay(ray core.Ray, params core.Interval) optional.Optional[Hit] {
	if !ray.Hits(bhv.boundingBox, params) {
		return optional.Empty[Hit]()
	}

	leftHit := bhv.leftChild.TestRay(ray, params)
	rightHit := bhv.rightChild.TestRay(ray, params)

	if leftHit.Present() && rightHit.Present() {
		return core.IfElse(leftHit.Value().Param < rightHit.Value().Param, leftHit, rightHit)
	}

	if leftHit.Present() {
		return leftHit
	}

	if rightHit.Present() {
		return rightHit
	}

	return optional.Empty[Hit]()
}

func (bvh *BVHNode) BoundingBox() core.Box {
	return bvh.boundingBox
}
