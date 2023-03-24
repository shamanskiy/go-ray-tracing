package objects

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type HitRecord struct {
	Param  core.Real
	Point  core.Vec3
	Normal core.Vec3
}

type Object interface {
	Hit(ray core.Ray) *HitRecord
	HitWithMin(ray core.Ray, minParam core.Real) *HitRecord
}
