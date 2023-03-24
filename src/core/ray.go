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
