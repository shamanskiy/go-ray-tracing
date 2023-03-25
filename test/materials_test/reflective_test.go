package materials

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/src/materials"
	"github.com/stretchr/testify/assert"
)

func TestReflective_ShouldReflectRayAroundNormal_WhenNotFuzzy(t *testing.T) {
	material := materials.NewReflective(anyColor, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(4, -3, 0)
	normalAtHitPoint := core.NewVec3(0, 1, 0)

	reflection := material.Reflect(incidentDirection, hitPoint, normalAtHitPoint)

	expected := materials.Reflection{
		Ray:         core.NewRay(hitPoint, core.NewVec3(4, 3, 0).Normalize()),
		Attenuation: material.Color(),
	}
	assert.Equal(t, expected, *reflection)
}

func TestReflective_ShouldNotReflect_WhenRayParallelToSurface(t *testing.T) {
	material := materials.NewReflective(anyColor, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(4, 0, 0)
	normalAtHitPoint := core.NewVec3(0, 1, 0)

	reflection := material.Reflect(incidentDirection, hitPoint, normalAtHitPoint)

	assert.Nil(t, reflection)
}

func TestReflective_ShouldNotReflect_WhenRaysNormalComponentCoalignedWithNormal(t *testing.T) {
	material := materials.NewReflective(anyColor, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(4, 3, 0)
	normalAtHitPoint := core.NewVec3(0, 1, 0)

	reflection := material.Reflect(incidentDirection, hitPoint, normalAtHitPoint)

	assert.Nil(t, reflection)
}

func TestRefective_WithFuzziness(t *testing.T) {
	material := materials.NewReflectiveFuzzy(anyColor, 0.5, random.NewRandomGenerator())
	incidentDirection := core.NewVec3(4, -3, 0)
	normalAtHitPoint := core.NewVec3(0, 1, 0)

	reflection := material.Reflect(incidentDirection, hitPoint, normalAtHitPoint)

	expectedMeanDirection := core.NewVec3(4, 3, 0).Normalize()
	randomPerturbation := reflection.Ray.Direction().Sub(expectedMeanDirection).Len()
	assert.Less(t, randomPerturbation, material.Fuzziness())
}

func TestReflective_FuzzinessCantBeLessThanZero(t *testing.T) {
	assert.Panics(t, func() {
		materials.NewReflectiveFuzzy(anyColor, -0.5, random.NewRandomGenerator())
	})
}

func TestReflective_FuzzinessCantBeMoreThanOne(t *testing.T) {
	assert.Panics(t, func() {
		materials.NewReflectiveFuzzy(anyColor, 1.5, random.NewRandomGenerator())
	})
}
