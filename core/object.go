package core

type Object interface {
	Hit(ray Ray) HitRecord
}
