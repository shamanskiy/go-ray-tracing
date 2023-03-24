package core

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (ray Ray) Eval(t Real) Vec3 {
	return ray.Origin.Add(ray.Direction.Mul(t))
}
