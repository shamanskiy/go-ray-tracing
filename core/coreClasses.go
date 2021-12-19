package core

import (
	"testing"

	"github.com/chewxy/math32"
	"github.com/go-gl/mathgl/mgl32"
)

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

func CheckVec3SameTol(t *testing.T, name string, result Vec3, expected Vec3) {
	if (result.Sub(expected).LenSqr()) < 1e-10 {
		t.Logf("\t\tPASSED: %v %v, expected %v", name, result, expected)
	} else {
		t.Fatalf("\t\tFAILED: %v %v, expected %v", name, result, expected)
	}
}

func CheckFloatSameTol(t *testing.T, name string, result Real, expected Real) {
	if math32.Abs(result-expected) < 1e-5 {
		t.Logf("\t\tPASSED: %v %v, expected %v", name, result, expected)
	} else {
		t.Fatalf("\t\tFAILED: %v %v, expected %v", name, result, expected)
	}
}
