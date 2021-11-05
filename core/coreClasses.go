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

func MulElem(a, b mgl32.Vec3) mgl32.Vec3 {
	return mgl32.Vec3{a.X() * b.X(), a.Y() * b.Y(), a.Z() * b.Z()}
}

var Red = Color{1.0, 0.0, 0.0}
var Green = Color{0.0, 1.0, 0.0}
var Blue = Color{0.0, 0.0, 1.0}
var Black = Color{0.0, 0.0, 0.0}
var White = Color{1.0, 1.0, 1.0}
var SkyBlue = Color{0.5, 0.7, 1.0}
var GrayMedium = Color{0.5, 0.5, 0.5}
