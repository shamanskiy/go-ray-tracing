package camera

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/chewxy/math32"
)

var globalUp = core.NewVec3(0, 1, 0)

type RayGenerator struct {
	origin          core.Vec3
	upperLeftCorner core.Vec3
	horizontalSpan  core.Vec3
	verticalSpan    core.Vec3

	right core.Vec3
	up    core.Vec3

	randomizer          random.RandomGenerator
	defocusBlurStrength core.Real
}

func NewRayGenerator(settings *CameraSettings, randomizer random.RandomGenerator) *RayGenerator {
	verticalFOVInRadians := settings.VerticalFOV * math32.Pi / 180

	halfHeight := math32.Tan(verticalFOVInRadians / 2)
	halfWidth := settings.AspectRatio * halfHeight

	back := settings.LookFrom.Sub(settings.LookAt).Normalize()
	right := globalUp.Cross(back).Normalize()
	up := back.Cross(right)

	focusDistance := settings.LookFrom.Sub(settings.LookAt).Len()
	originToCorner := up.Mul(halfHeight).Sub(right.Mul(halfWidth)).Sub(back).Mul(focusDistance)

	return &RayGenerator{
		origin:              settings.LookFrom,
		upperLeftCorner:     settings.LookFrom.Add(originToCorner),
		horizontalSpan:      right.Mul(2 * halfWidth * focusDistance),
		verticalSpan:        up.Mul(-2 * halfHeight * focusDistance),
		right:               right,
		up:                  up,
		randomizer:          randomizer,
		defocusBlurStrength: settings.DefocusBlurStrength,
	}
}

func (r *RayGenerator) GenerateRay(u, v core.Real) core.Ray {
	focusPlanePoint := r.upperLeftCorner.Add(r.horizontalSpan.Mul(u)).Add(r.verticalSpan.Mul(v))
	cameraOrigin := r.origin
	if r.defocusBlurStrength > 0 {
		cameraOrigin = cameraOrigin.Add(r.randomOriginOffset())
	}
	rayDirection := focusPlanePoint.Sub(cameraOrigin)

	return core.NewRay(cameraOrigin, rayDirection)
}

func (r *RayGenerator) randomOriginOffset() core.Vec3 {
	randomVec2 := r.randomizer.Vec3InUnitDisk()
	return r.right.Mul(randomVec2.X()).Add(r.up.Mul(randomVec2.Y())).Mul(r.defocusBlurStrength)
}
