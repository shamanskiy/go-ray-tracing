package core

import (
	"testing"

	"github.com/go-gl/mathgl/mgl32"
)

func TestRay(t *testing.T) {
	ray := Ray{mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{1.0, 2.0, 3.0}}
	t.Logf("Given a ray with origin at %v and direction %v,\n", ray.Origin, ray.Direction)

	t.Log("\twe can evaluate the ray at point t = 2.0:")
	point := ray.Eval(2.0)
	expected := mgl32.Vec3{2.0, 4.0, 6.0}
	if point == expected {
		t.Logf("\t\tPASSED: result is %v, expected %v.\n", point, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v.\n", point, expected)
	}

	t.Log("\twe can evaluate the ray at point t = -2.0:")
	point = ray.Eval(-2.0)
	expected = mgl32.Vec3{-2.0, -4.0, -6.0}
	if point == expected {
		t.Logf("\t\tPASSED: result is %v, expected %v.\n", point, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v.\n", point, expected)
	}
}
