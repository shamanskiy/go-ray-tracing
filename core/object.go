package core

type Object interface {
	Hit(ray Ray) HitRecord
	HitWithMin(ray Ray, minParam Real) HitRecord
}
