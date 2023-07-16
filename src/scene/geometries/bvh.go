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

func BuildBVH(hittables []Hittable) *BVHNode {
	sortingAxis := rand.Intn(3)
	sort.Slice(hittables, func(i, j int) bool {
		return compareGeometries(hittables, i, j, sortingAxis)
	})

	bvhNode := &BVHNode{}

	switch len(hittables) {
	case 0:
		bvhNode.leftChild = emptyHittable{}
		bvhNode.rightChild = emptyHittable{}
	case 1:
		bvhNode.leftChild = hittables[0]
		bvhNode.rightChild = emptyHittable{}
	default:
		bvhNode.leftChild = BuildBVH(hittables[0 : len(hittables)/2])
		bvhNode.rightChild = BuildBVH(hittables[len(hittables)/2:])
	}

	bvhNode.boundingBox = bvhNode.leftChild.BoundingBox().Union(bvhNode.rightChild.BoundingBox())

	return bvhNode
}

func compareGeometries(hittables []Hittable, i, j, sortingAxis int) bool {
	minI := hittables[i].BoundingBox().Min().At(sortingAxis)
	minj := hittables[j].BoundingBox().Min().At(sortingAxis)
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
