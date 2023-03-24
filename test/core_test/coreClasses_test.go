package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/test"
)

func TestRay(t *testing.T) {
	ray := core.Ray{core.Vec3{0.0, 0.0, 0.0}, core.Vec3{1.0, 2.0, 3.0}}
	t.Logf("Given a ray with origin at %v and direction %v,\n", ray.Origin, ray.Direction)

	t.Log("\twe can evaluate the ray at point t = 2.0:")
	point := ray.Eval(2.0)
	expected := core.Vec3{2.0, 4.0, 6.0}
	test.CheckResult(t, "vector", point, expected)

	t.Log("\twe can evaluate the ray at point t = -2.0:")
	point = ray.Eval(-2.0)
	expected = core.Vec3{-2.0, -4.0, -6.0}
	test.CheckResult(t, "vector", point, expected)
}

func TestMulElem_Vec3(t *testing.T) {
	vecA, vecB := core.Vec3{1., 2., 3.}, core.Vec3{4., 5., 6.}
	t.Logf("Given two vectors %v and %v\n,", vecA, vecB)

	t.Log("\twe can multiply them element-wise:")
	prod := core.MulElem(vecA, vecB)
	expected := core.Vec3{4., 10., 18}

	test.CheckResult(t, "product", prod, expected)
}

func TestMulElem_Color(t *testing.T) {
	colorA, colorB := core.White, core.Red
	t.Logf("Given two colors %v and %v\n,", colorA, colorB)

	t.Log("\twe can multiply them element-wise:")
	prod := core.MulElem(colorA, colorB)
	expected := core.Red

	test.CheckResult(t, "product", prod, expected)
}

func TestReflect(t *testing.T) {
	vector := core.Vec3{3., -4, 0.}
	axis := core.Vec3{0., 1., 0.}
	t.Logf("Given a vector %v and an axis %v,\n", vector, axis)

	t.Log("\twe can reflect the vector around the axis:")
	reflected := core.Reflect(vector, axis)
	expected := core.Vec3{3., 4., 0.}

	test.CheckResult(t, "reflected vector", reflected, expected)
}

func TestIsSameReal(t *testing.T) {
	A := core.Real(1.0)
	B := core.Real(1.000001)
	t.Logf("Given two real numbers %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", core.RealTolerance())
	test.CheckResult(t, "Numbers are within tolerance", core.IsSameReal(A, B), true)

	A = core.Real(1.0)
	B = core.Real(2.0)
	t.Logf("Given two real numbers %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", core.RealTolerance())
	test.CheckResult(t, "Numbers are within tolerance", core.IsSameReal(A, B), false)

	A = core.Real(1.0)
	B = core.Real(1.00001)
	t.Logf("Given two real numbers %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", core.RealTolerance())
	test.CheckResult(t, "Numbers are within tolerance", core.IsSameReal(A, B), false)
}

func TestIsSameVec3(t *testing.T) {
	A := core.Vec3{1.0, 0.0, 0.0}
	B := core.Vec3{1.000001, 0.0, 0.0}
	t.Logf("Given two vectors %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", core.RealTolerance())
	test.CheckResult(t, "Numbers are within tolerance", core.IsSameVec3(A, B), true)

	A = core.Vec3{1.0, 0.0, 0.0}
	B = core.Vec3{2.0, 0.0, 0.0}
	t.Logf("Given two vectors %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", core.RealTolerance())
	test.CheckResult(t, "Numbers are within tolerance", core.IsSameVec3(A, B), false)

	A = core.Vec3{1.0, 0.0, 0.0}
	B = core.Vec3{1.00001, 0.0, 0.0}
	t.Logf("Given two vectors %v and %v,\n", A, B)
	t.Logf("  we can check if they are within tolerance %v:", core.RealTolerance())
	test.CheckResult(t, "Numbers are within tolerance", core.IsSameVec3(A, B), false)

}

func TestDiv(t *testing.T) {
	A := core.Vec3{1., 2., 3.}
	b := core.Real(2.)
	t.Logf("Given a vector %v and a scalar %v,\n", A, b)
	t.Log("  we can divide the vector by the scalar:")
	test.CheckResult(t, "Division result", core.Div(A, b), core.Vec3{0.5, 1., 1.5})
}
