package objects

import "github.com/Shamanskiy/go-ray-tracer/core"

type Object interface {
	Hit(ray core.Ray) core.HitRecord
	HitWithMin(ray core.Ray, minParam core.Real) core.HitRecord
}
