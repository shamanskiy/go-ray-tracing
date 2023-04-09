package materials_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/scene/materials"
	"github.com/stretchr/testify/assert"
)

func TestDiffusiveLight_ShouldEmitLight(t *testing.T) {
	material := materials.NewDiffusiveLight(MATERIAL_COLOR)

	reflection := material.Reflect(RAY_DIRECTION, HIT_POINT, NORMAL_AT_HIT_POINT)

	assert.Equal(t, materials.Emitted, reflection.Type)
	assert.Equal(t, MATERIAL_COLOR, reflection.Color)
}
