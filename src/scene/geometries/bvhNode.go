package geometries

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type BVHNode struct {
	boundingBox core.Box
	leftChild   Hittable
	rightChild  Hittable
}

// func (bhv *BVHNode) TestRay(ray core.Ray) (hitParams []core.Real) {
// 	if !ray.Hits(bhv.boundingBox) {
// 		return []core.Real{}
// 	}
// }
