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

func (ray Ray) Hits(box Box, params Interval) bool {
	for i := 0; i < 3; i++ {
		invD := 1. / ray.direction.At(i)
		t0 := (box.min.At(i) - ray.origin.At(i)) * invD
		t1 := (box.max.At(i) - ray.origin.At(i)) * invD
		if invD < 0 {
			t0, t1 = t1, t0
		}

		params.min = ternaryIf(t0 > params.min, t0, params.min)
		params.max = ternaryIf(t1 < params.max, t1, params.max)

		if params.max < params.min {
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
