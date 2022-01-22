package core

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestRay(t *testing.T) {
	ray := Ray{Vec3{0.0, 0.0, 0.0}, Vec3{1.0, 2.0, 3.0}}
	t.Logf("Given a ray with origin at %v and direction %v,\n", ray.Origin, ray.Direction)

	t.Log("\twe can evaluate the ray at point t = 2.0:")
	point := ray.Eval(2.0)
	expected := Vec3{2.0, 4.0, 6.0}
	utils.CheckResult(t, "vector", point, expected)

	t.Log("\twe can evaluate the ray at point t = -2.0:")
	point = ray.Eval(-2.0)
	expected = Vec3{-2.0, -4.0, -6.0}
	utils.CheckResult(t, "vector", point, expected)
}

func TestMulElem_Vec3(t *testing.T) {
	vecA, vecB := Vec3{1., 2., 3.}, Vec3{4., 5., 6.}
	t.Logf("Given two vectors %v and %v\n,", vecA, vecB)

	t.Log("\twe can multiply them element-wise:")
	prod := MulElem(vecA, vecB)
	expected := Vec3{4., 10., 18}

	utils.CheckResult(t, "product", prod, expected)
}

func TestMulElem_Color(t *testing.T) {
	colorA, colorB := White, Red
	t.Logf("Given two colors %v and %v\n,", colorA, colorB)

	t.Log("\twe can multiply them element-wise:")
	prod := MulElem(colorA, colorB)
	expected := Red

	utils.CheckResult(t, "product", prod, expected)
}

func TestReflect(t *testing.T) {
	vector := Vec3{3., -4, 0.}
	axis := Vec3{0., 1., 0.}
	t.Logf("Given a vector %v and an axis %v,\n", vector, axis)

	t.Log("\twe can reflect the vector around the axis:")
	reflected := Reflect(vector, axis)
	expected := Vec3{3., 4., 0.}

	utils.CheckResult(t, "reflected vector", reflected, expected)
}

func TestIsSameReal(t *testing.T) {
	A := Real(1.0)
	B := Real(1.000001)
	t.Logf("Given two real numbers %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", RealTolerance())
	utils.CheckResult(t, "Numbers are within tolerance", IsSameReal(A, B), true)

	A = Real(1.0)
	B = Real(2.0)
	t.Logf("Given two real numbers %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", RealTolerance())
	utils.CheckResult(t, "Numbers are within tolerance", IsSameReal(A, B), false)

	A = Real(1.0)
	B = Real(1.00001)
	t.Logf("Given two real numbers %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", RealTolerance())
	utils.CheckResult(t, "Numbers are within tolerance", IsSameReal(A, B), false)
}

func TestIsSameVec3(t *testing.T) {
	A := Vec3{1.0, 0.0, 0.0}
	B := Vec3{1.000001, 0.0, 0.0}
	t.Logf("Given two vectors %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", RealTolerance())
	utils.CheckResult(t, "Numbers are within tolerance", IsSameVec3(A, B), true)

	A = Vec3{1.0, 0.0, 0.0}
	B = Vec3{2.0, 0.0, 0.0}
	t.Logf("Given two vectors %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", RealTolerance())
	utils.CheckResult(t, "Numbers are within tolerance", IsSameVec3(A, B), false)

	A = Vec3{1.0, 0.0, 0.0}
	B = Vec3{1.00001, 0.0, 0.0}
	t.Logf("Given two vectors %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", RealTolerance())
	utils.CheckResult(t, "Numbers are within tolerance", IsSameVec3(A, B), false)

}

func TestDiv(t *testing.T) {
	A := Vec3{1., 2., 3.}
	b := Real(2.)
	t.Logf("Given a vector %v and a scalar %v,\n", A, b)
	t.Log("  we can divide the vector by the scalar:")
	utils.CheckResult(t, "Division result", Div(A, b), Vec3{0.5, 1., 1.5})
}
