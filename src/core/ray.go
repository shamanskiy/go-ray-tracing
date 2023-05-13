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

// Duplicating the code for X,Y and Z for better performance.
// I tried using vec3.At but it turned out to be too slow because of the switch.
func (ray Ray) Hits(box Box, params Interval) bool {
	invD := 1. / ray.direction.X()
	t0 := (box.min.X() - ray.origin.X()) * invD
	t1 := (box.max.X() - ray.origin.X()) * invD
	if invD < 0 {
		t0, t1 = t1, t0
	}

	params.min = IfElse(t0 > params.min, t0, params.min)
	params.max = IfElse(t1 < params.max, t1, params.max)

	if params.max < params.min {
		return false
	}
	invD = 1. / ray.direction.Y()
	t0 = (box.min.Y() - ray.origin.Y()) * invD
	t1 = (box.max.Y() - ray.origin.Y()) * invD
	if invD < 0 {
		t0, t1 = t1, t0
	}

	params.min = IfElse(t0 > params.min, t0, params.min)
	params.max = IfElse(t1 < params.max, t1, params.max)

	if params.max < params.min {
		return false
	}
	invD = 1. / ray.direction.Z()
	t0 = (box.min.Z() - ray.origin.Z()) * invD
	t1 = (box.max.Z() - ray.origin.Z()) * invD
	if invD < 0 {
		t0, t1 = t1, t0
	}

	params.min = IfElse(t0 > params.min, t0, params.min)
	params.max = IfElse(t1 < params.max, t1, params.max)

	if params.max < params.min {
		return false
	}
	return true
}
