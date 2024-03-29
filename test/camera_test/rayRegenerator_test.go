package camera_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/camera"
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/stretchr/testify/assert"
)

var LOOK_FROM = core.NewVec3(0, 0, 0)
var LOOK_AT = core.NewVec3(0, 0, -2)
var VERTICAL_FOV = core.Real(90.)
var ASPECT_RATIO = core.Real(2.)
var CAMERA_SETTINGS = &camera.CameraSettings{
	VerticalFOV: VERTICAL_FOV,
	AspectRatio: ASPECT_RATIO,
	LookAt:      LOOK_AT,
	LookFrom:    LOOK_FROM,
}

func TestRayGenerator_ShouldGenerateRays(t *testing.T) {
	rayGenerator := camera.NewRayGenerator(CAMERA_SETTINGS, random.NewRandomGenerator())

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

func TestRayGenerator_RayWithDefocusBlurShouldFocusOnFocusPlane(t *testing.T) {
	settings := *CAMERA_SETTINGS
	settings.DefocusBlurStrength = 1
	rayGenerator := camera.NewRayGenerator(&settings, random.NewRandomGenerator())

	for i := 0; i < 10; i++ {
		ray := rayGenerator.GenerateRay(0.5, 0.5)
		assert.Equal(t, LOOK_AT, ray.Eval(1))
		assert.NotEqual(t, LOOK_FROM, ray.Origin())
	}
}
