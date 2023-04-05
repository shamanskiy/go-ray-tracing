package camera_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/camera"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

var LOOK_FROM = core.NewVec3(0, 0, 0)
var LOOK_AT = core.NewVec3(0, 0, -2)
var VERTICAL_FOV = core.Real(90.)
var ASPECT_RATIO = core.Real(2.)

func TestRayGenerator_ShouldGenerateRays(t *testing.T) {
	rayGenerator := camera.NewRayGenerator(LOOK_FROM, LOOK_AT, VERTICAL_FOV, ASPECT_RATIO)

	centerRay := rayGenerator.GenerateRay(0.5, 0.5)
	assert.Equal(t, LOOK_FROM, centerRay.Origin())
	assert.Equal(t, LOOK_AT, centerRay.Eval(1))

	topLeftRay := rayGenerator.GenerateRay(0, 0)
	assert.Equal(t, core.NewVec3(-4, 2, -2), topLeftRay.Eval(1))

	topRightRay := rayGenerator.GenerateRay(1, 0)
	assert.Equal(t, core.NewVec3(4, 2, -2), topRightRay.Eval(1))

	bottomLeftRay := rayGenerator.GenerateRay(0, 1)
	assert.Equal(t, core.NewVec3(-4, -2, -2), bottomLeftRay.Eval(1))

	bottomRightRay := rayGenerator.GenerateRay(1, 1)
	assert.Equal(t, core.NewVec3(4, -2, -2), bottomRightRay.Eval(1))
}
