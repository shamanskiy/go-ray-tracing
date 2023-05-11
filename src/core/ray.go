package core

type Ray struct {
	origin    Vec3
	direction Vec3
}

func NewRay(origin Vec3, direction Vec3) Ray {
	return Ray{origin, direction}
}

func (ray Ray) Origin() Vec3 {
	return ray.origin
}

func (ray Ray) Direction() Vec3 {
	return ray.direction
}

func (ray Ray) Eval(t Real) Vec3 {
	return ray.origin.Add(ray.direction.Mul(t))
}

func (ray Ray) Hits(box Box, tMin, tMax Real) bool {
	for i := 0; i < 3; i++ {
		invD := 1. / ray.direction.At(i)
		t0 := (box.min.At(i) - ray.origin.At(i)) * invD
		t1 := (box.max.At(i) - ray.origin.At(i)) * invD
		if invD < 0 {
			t0, t1 = t1, t0
		}

		tMin = ternaryIf(t0 > tMin, t0, tMin)
		tMax = ternaryIf(t1 < tMax, t1, tMax)

		if tMax < tMin {
			return false
		}
	}
	return true
}

func ternaryIf[T any](condition bool, ifTrue, ifFalse T) T {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}
