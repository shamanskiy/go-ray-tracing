package core

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Vec3 struct {
	vec mgl32.Vec3
}

func NewVec3(x, y, z Real) Vec3 {
	return Vec3{vec: mgl32.Vec3{x, y, z}}
}

func (vec Vec3) X() Real {
	return vec.vec.X()
}

func (vec Vec3) Y() Real {
	return vec.vec.Y()
}

func (vec Vec3) Z() Real {
	return vec.vec.Z()
}

func (vecA Vec3) Add(vecB Vec3) Vec3 {
	return Vec3{vecA.vec.Add(vecB.vec)}
}

func (vecA Vec3) Sub(vecB Vec3) Vec3 {
	return Vec3{vecA.vec.Sub(vecB.vec)}
}

func (vec Vec3) Mul(scalar Real) Vec3 {
	return Vec3{vec.vec.Mul(scalar)}
}

func (vecA Vec3) MulVec(vecB Vec3) Vec3 {
	return NewVec3(vecA.X()*vecB.X(), vecA.Y()*vecB.Y(), vecA.Z()*vecB.Z())
}

func (vecA Vec3) Cross(vecB Vec3) Vec3 {
	return Vec3{vecA.vec.Cross(vecB.vec)}
}

func (vec Vec3) Div(scalar Real) Vec3 {
	return NewVec3(vec.X()/scalar, vec.Y()/scalar, vec.Z()/scalar)
}

func (vecA Vec3) Dot(vecB Vec3) Real {
	return vecA.vec.Dot(vecB.vec)
}

func (vec Vec3) Len() Real {
	return vec.vec.Len()
}

func (vec Vec3) LenSqr() Real {
	return vec.vec.LenSqr()
}

func (vec Vec3) Normalize() Vec3 {
	return Vec3{vec.vec.Normalize()}
}

func (vec Vec3) Reflect(axis Vec3) Vec3 {
	return vec.Sub(axis.Mul(2 * vec.Dot(axis)))
}

func (vecA Vec3) InDelta(vecB Vec3, delta Real) bool {
	return vecA.Sub(vecB).LenSqr() < delta*delta
}
