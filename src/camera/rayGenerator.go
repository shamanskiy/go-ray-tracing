package camera

import (
	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/chewxy/math32"
)

var globalUp = core.NewVec3(0, 1, 0)

type RayGenerator struct {
	origin          core.Vec3
	upperLeftCorner core.Vec3
	horizontalSpan  core.Vec3
	verticalSpan    core.Vec3
}

func NewRayGenerator(lookFrom, lookAt core.Vec3, verticalFOV, aspectRatio core.Real) *RayGenerator {
	verticalFOVInRadians := verticalFOV * math32.Pi / 180

	halfHeight := math32.Tan(verticalFOVInRadians / 2)
	halfWidth := aspectRatio * halfHeight

	back := lookFrom.Sub(lookAt).Normalize()
	right := globalUp.Cross(back).Normalize()
	up := back.Cross(right)

	focusDistance := lookFrom.Sub(lookAt).Len()
	originToCorner := up.Mul(halfHeight).Sub(right.Mul(halfWidth)).Sub(back).Mul(focusDistance)

	return &RayGenerator{
		origin:          lookFrom,
		upperLeftCorner: lookFrom.Add(originToCorner),
		horizontalSpan:  right.Mul(2 * halfWidth * focusDistance),
		verticalSpan:    up.Mul(-2 * halfHeight * focusDistance),
	}
}

func (r *RayGenerator) GenerateRay(u, v core.Real) core.Ray {
	rayDirection := r.upperLeftCorner.Add(r.horizontalSpan.Mul(u)).Add(r.verticalSpan.Mul(v)).Sub(r.origin)
	ray := core.NewRay(r.origin, rayDirection)
	return ray
}
