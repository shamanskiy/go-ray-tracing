package geometries

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type BVHNode struct {
	boundingBox core.Box
	leftChild   Hittable
	rightChild  Hittable
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
