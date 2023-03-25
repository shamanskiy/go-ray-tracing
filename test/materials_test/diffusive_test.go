package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/color"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/Shamanskiy/go-ray-tracer/src/objects"
	"github.com/stretchr/testify/assert"
)

var ray = core.NewRay(core.NewVec3(1, 2, 3), core.NewVec3(4, 5, 6))
var hitRecord = objects.HitRecord{
	Param:  1,
	Point:  core.NewVec3(0, 1, 2),
	Normal: core.NewVec3(0, 0, 1),
}

func TestDiffusive_ShouldReflectRayInNormalDirection_WhenNotRandom(t *testing.T) {
	material := materials.NewDiffusive(color.Red, random.NewFakeRandomGenerator())

	reflection := material.Reflect(ray, hitRecord)

	expected := materials.Reflection{
		Ray:         core.NewRay(hitRecord.Point, hitRecord.Normal),
		Attenuation: material.Color(),
	}
	assert.Equal(t, expected, *reflection)
}

func TestDiffusive_ShouldReflectRayWithinUnitSphereOfNormal_WhenRandom(t *testing.T) {
	material := materials.NewDiffusive(color.Red, random.NewRandomGenerator())

	reflection := material.Reflect(ray, hitRecord)

	randomPerturbation := reflection.Ray.Direction().Sub(hitRecord.Normal).Len()
	assert.Less(t, randomPerturbation, core.Real(1))
	assert.Equal(t, material.Color(), reflection.Attenuation)
	assert.Equal(t, hitRecord.Point, reflection.Ray.Origin())
}
