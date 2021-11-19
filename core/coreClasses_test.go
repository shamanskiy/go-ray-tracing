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
