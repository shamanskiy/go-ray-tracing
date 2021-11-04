package core

import "github.com/go-gl/mathgl/mgl32"

type Color = mgl32.Vec3
type Vec3 = mgl32.Vec3
type Real = float32

type Ray struct {
	Origin    Vec3
	Direction Vec3
}

func (ray Ray) Eval(t float32) Vec3 {
	return ray.Origin.Add(ray.Direction.Mul(t))
}

func MulElem(a, b Color) Color {
	return Color{a.X() * b.X(), a.Y() * b.Y(), a.Z() * b.Z()}
}
