package core

import "github.com/go-gl/mathgl/mgl32"

type Color = mgl32.Vec3
type Vec3 = mgl32.Vec3
type Real = float32

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (ray Ray) Eval(t Real) Vec3 {
	return ray.Origin.Add(ray.Direction.Mul(t))
}

func MulElem(a, b Vec3) Vec3 {
	return Vec3{a.X() * b.X(), a.Y() * b.Y(), a.Z() * b.Z()}
}

func Reflect(vec Vec3, axis Vec3) Vec3 {
	return vec.Sub(axis.Mul(2 * vec.Dot(axis)))
}
