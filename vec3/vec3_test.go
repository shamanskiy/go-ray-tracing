package vec3

import "testing"

// Test addition of 3D vectors
func TestAdd(t *testing.T) {
	A := Vec3{1.0, 2.0, 3.0}
	B := Vec3{4.0, 5.0, 6.0}
	t.Logf("Given two 3D vectors %v and %v,", A, B)

	t.Log("\twe can add them:")
	sum := A.Add(B)
	sum_expected := Vec3{5.0, 7.0, 9.0}
	if sum == sum_expected {
		t.Logf("\t\tPASSED: %v", sum)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", sum, sum_expected)
	}

	t.Log("\twe can subtract one from another:")
	diff := A.Sub(B)
	diff_expected := Vec3{-3.0, -3.0, -3.0}
	if diff == diff_expected {
		t.Logf("\t\tPASSED: %v", diff)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", diff, diff_expected)
	}

}
