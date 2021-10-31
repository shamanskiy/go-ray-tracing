package core

import (
	"testing"
)

func TestRay(t *testing.T) {
	ray := Ray{Origin: Vec3{0.0, 0.0, 0.0}, Direction: Vec3{1.0, 2.0, 3.0}}
	t.Logf("Given a ray with origin at %v and direction %v,\n", ray.Origin, ray.Direction)

	t.Log("\twe can evaluate the ray at point t = 2.0:")
	point := ray.Eval(2.0)
	expected := Vec3{2.0, 4.0, 6.0}
	if point == expected {
		t.Logf("\t\tPASSED: result is %v, expected %v.\n", point, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v.\n", point, expected)
	}

	t.Log("\twe can evaluate the ray at point t = -2.0:")
	point = ray.Eval(-2.0)
	expected = Vec3{-2.0, -4.0, -6.0}
	if point == expected {
		t.Logf("\t\tPASSED: result is %v, expected %v.\n", point, expected)
	} else {
		t.Fatalf("\t\tFAILED: result is %v, expected %v.\n", point, expected)
	}
}
