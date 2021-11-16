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
