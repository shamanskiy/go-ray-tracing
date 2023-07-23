package materials_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/stretchr/testify/assert"
)

const LIGHT_INTENSITY = 2

func TestDiffusiveLight_ShouldEmitLight(t *testing.T) {
	material := materials.NewDiffusiveLight(MATERIAL_COLOR, LIGHT_INTENSITY)

	reflection := material.Reflect(RAY_DIRECTION, HIT_POINT, NORMAL_AT_HIT_POINT)

	assert.Equal(t, materials.Emitted, reflection.Type)
	assert.Equal(t, MATERIAL_COLOR.Mul(LIGHT_INTENSITY), reflection.Color)
}
