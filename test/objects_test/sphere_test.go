package objects

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/stretchr/testify/assert"
)

func TestSphere_ShouldReturnTwoDistinctHits_IfRayIntersectsSphere(t *testing.T) {
	sphere := objects.Sphere{core.NewVec3(0, 0, 0), 2.0}
	ray := core.NewRay(core.NewVec3(4, 0, 0), core.NewVec3(-1, 0, 0))

	hits := sphere.TestRay(ray)

	assert.Equal(t, []core.Real{2, 6}, hits)
}

func TestSphere_ShouldReturnOneHitTwice_IfRayTouchesSphere(t *testing.T) {
	sphere := objects.Sphere{core.NewVec3(0, 0, 0), 2.0}
	ray := core.NewRay(core.NewVec3(4, 2, 0), core.NewVec3(-1, 0, 0))

	hits := sphere.TestRay(ray)

	assert.Equal(t, []core.Real{4, 4}, hits)
}

func TestSphere_ShouldReturnNoHits_IfRayDoesNotIntersectSphere(t *testing.T) {
	sphere := objects.Sphere{core.NewVec3(0, 0, 0), 2.0}
	ray := core.NewRay(core.NewVec3(4, 4, 0), core.NewVec3(-1, 0, 0))

	hits := sphere.TestRay(ray)

	assert.Empty(t, hits)
}

func TestSphere_ShouldEvaluateHit(t *testing.T) {
	sphere := objects.Sphere{core.NewVec3(0, 0, 0), 2.0}
	ray := core.NewRay(core.NewVec3(4, 0, 0), core.NewVec3(-1, 0, 0))

	hitRecord := sphere.EvaluateHit(ray, 2)

	assert.Equal(t, core.NewVec3(2, 0, 0), hitRecord.Point)
	assert.Equal(t, core.NewVec3(1, 0, 0), hitRecord.Normal)
}
