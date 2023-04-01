package objects

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type Plane struct {
	origin core.Vec3
	normal core.Vec3
}

func NewPlane(origin, normal core.Vec3) Plane {
	return Plane{
		origin: origin,
		normal: normal,
	}
}

func (p Plane) TestRay(ray core.Ray) (hitParams []core.Real) {
	return []core.Real{}
}

func (p Plane) EvaluateHit(ray core.Ray, hitParam core.Real) HitRecord {
	return HitRecord{}
}
