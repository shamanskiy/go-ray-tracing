package materials_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/stretchr/testify/assert"
)

func TestReflective_ShouldReflectRayAroundNormal_WhenNotFuzzy(t *testing.T) {
	material := materials.NewReflective(MATERIAL_COLOR, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(4, -3, 0)

	reflection := material.Reflect(incidentDirection, HIT_POINT, NORMAL_AT_HIT_POINT)

	expected := materials.Reflection{
		Type:  materials.Scattered,
		Ray:   core.NewRay(HIT_POINT, core.NewVec3(4, 3, 0).Normalize()),
		Color: MATERIAL_COLOR,
	}
	assert.Equal(t, expected, reflection)
}

func TestReflective_ShouldNotReflect_WhenRayParallelToSurface(t *testing.T) {
	material := materials.NewReflective(MATERIAL_COLOR, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(4, 0, 0)

	reflection := material.Reflect(incidentDirection, HIT_POINT, NORMAL_AT_HIT_POINT)

	assert.Equal(t, materials.Absorbed, reflection.Type)
}

func TestReflective_ShouldNotReflect_WhenRaysNormalComponentCoalignedWithNormal(t *testing.T) {
	material := materials.NewReflective(MATERIAL_COLOR, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(4, 3, 0)

	reflection := material.Reflect(incidentDirection, HIT_POINT, NORMAL_AT_HIT_POINT)

	assert.Equal(t, materials.Absorbed, reflection.Type)
}

func TestRefective_WithFuzziness(t *testing.T) {
	var fuzziness core.Real = 0.5
	material := materials.NewReflectiveFuzzy(MATERIAL_COLOR, fuzziness, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(4, -3, 0)

	reflection := material.Reflect(incidentDirection, HIT_POINT, NORMAL_AT_HIT_POINT)

	expectedMeanDirection := core.NewVec3(4, 3, 0).Normalize()
	randomPerturbation := reflection.Ray.Direction().Sub(expectedMeanDirection).Len()
	assert.Less(t, randomPerturbation, fuzziness)
}

func TestReflective_FuzzinessCantBeLessThanZero(t *testing.T) {
	assert.Panics(t, func() {
		materials.NewReflectiveFuzzy(MATERIAL_COLOR, -0.5, random.NewRandomGenerator())
	})
}

func TestReflective_FuzzinessCantBeMoreThanOne(t *testing.T) {
	assert.Panics(t, func() {
		materials.NewReflectiveFuzzy(MATERIAL_COLOR, 1.5, random.NewRandomGenerator())
	})
}
