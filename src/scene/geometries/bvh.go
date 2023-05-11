package geometries

import (
	"math/rand"
	"sort"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
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

func (bhv *BVHNode) TestRay(ray core.Ray, params core.Interval) Hit {
	if !ray.Hits(bhv.boundingBox, params) {
		return Hit{}
	}

	leftHit := bhv.leftChild.TestRay(ray, params)
	rightHit := bhv.rightChild.TestRay(ray, params)

	if leftHit.HasHit && rightHit.HasHit {
		return core.IfElse(leftHit.Param < rightHit.Param, leftHit, rightHit)
	}

	if leftHit.HasHit {
		return leftHit
	}

	if rightHit.HasHit {
		return rightHit
	}

	return Hit{}
}

func (bvh *BVHNode) BoundingBox() core.Box {
	return bvh.boundingBox
}
