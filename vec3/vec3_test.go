package vec3

import (
	"testing"
)

func TestVec3_Add(t *testing.T) {
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
}

func TestVec3_Diff(t *testing.T) {
	A := Vec3{1.0, 2.0, 3.0}
	B := Vec3{4.0, 5.0, 6.0}
	t.Logf("Given two 3D vectors %v and %v,", A, B)

	t.Log("\twe can subtract one from another:")
	diff := A.Sub(B)
	diff_expected := Vec3{-3.0, -3.0, -3.0}

	if diff == diff_expected {
		t.Logf("\t\tPASSED: %v", diff)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", diff, diff_expected)
	}
}

func TestVec3_MultVec(t *testing.T) {
	A := Vec3{1.0, 2.0, 3.0}
	B := Vec3{4.0, 5.0, 6.0}
	t.Logf("Given two 3D vectors %v and %v,", A, B)

	t.Log("\twe can multiply them element-wise:")
	prod := A.MultVec(B)
	prod_expected := Vec3{4.0, 10.0, 18.0}

	if prod == prod_expected {
		t.Logf("\t\tPASSED: %v", prod)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", prod, prod_expected)
	}
}

func TestVec3_Mult(t *testing.T) {
	A := Vec3{1.0, 2.0, 3.0}
	b := 2.0
	t.Logf("Given a 3D vector %v and a scalar %v,", A, b)

	t.Log("\twe can multiply them :")
	prod := A.Mult(b)
	prod_expected := Vec3{2.0, 4.0, 6.0}

	if prod == prod_expected {
		t.Logf("\t\tPASSED: %v", prod)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", prod, prod_expected)
	}
}

func TestVec3_DivVec(t *testing.T) {
	A := Vec3{1.0, 2.0, 3.0}
	B := Vec3{4.0, 5.0, 6.0}
	t.Logf("Given two 3D vectors %v and %v,", A, B)

	t.Log("\twe can divide one by another element-wise:")
	ratio := A.DivVec(B)
	ratio_expected := Vec3{0.25, 0.4, 0.5}

	if ratio == ratio_expected {
		t.Logf("\t\tPASSED: %v", ratio)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", ratio, ratio_expected)
	}
}

func TestVec3_Di(t *testing.T) {
	A := Vec3{1.0, 2.0, 3.0}
	b := 2.0
	t.Logf("Given a 3D vector %v and a scalar %v,", A, b)

	t.Log("\twe can divide the vector by the scalar:")
	ratio := A.Div(b)
	ratio_expected := Vec3{0.5, 1.0, 1.5}

	if ratio == ratio_expected {
		t.Logf("\t\tPASSED: %v", ratio)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", ratio, ratio_expected)
	}
}

func TestVec3_DotVec3(t *testing.T) {
	A := Vec3{1.0, 2.0, 3.0}
	B := Vec3{4.0, 5.0, 6.0}
	t.Logf("Given two 3D vectors %v and %v,", A, B)

	t.Log("\twe can compute their dot product:")
	prod := A.Dot(B)
	prod_expected := 32.0

	if prod == prod_expected {
		t.Logf("\t\tPASSED: %v", prod)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", prod, prod_expected)
	}
}

func TestVec3_CrossVec3(t *testing.T) {
	A := Vec3{1.0, 0.0, 0.0}
	B := Vec3{0.0, 1.0, 0.0}
	t.Logf("Given two 3D vectors %v and %v,", A, B)

	t.Log("\twe can compute their cross product:")
	prod := A.Cross(B)
	prod_expected := Vec3{0.0, 0.0, 1.0}

	if prod == prod_expected {
		t.Logf("\t\tPASSED: %v", prod)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", prod, prod_expected)
	}
}

func TestVec3_Length(t *testing.T) {
	A := Vec3{2.0, 3.0, 6.0}
	t.Logf("Given a 3D vector %v", A)

	t.Log("\twe can compute its length:")
	len := A.Length()
	len_expected := 7.0

	if len == len_expected {
		t.Logf("\t\tPASSED: %v", len)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", len, len_expected)
	}
}

func TestVec3_LengthSquared(t *testing.T) {
	A := Vec3{2.0, 3.0, 6.0}
	t.Logf("Given a 3D vector %v", A)

	t.Log("\twe can compute its squared length:")
	len := A.LengthSquared()
	len_expected := 49.0

	if len == len_expected {
		t.Logf("\t\tPASSED: %v", len)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", len, len_expected)
	}
}

func TestVec3_Normalize(t *testing.T) {
	A := Vec3{9.0, 12.0, 20.0}
	t.Logf("Given a 3D vector %v", A)

	t.Log("\twe can normalize it:")
	unit := A.Normalize()
	unit_expected := Vec3{0.36, 0.48, 0.8}

	if unit == unit_expected {
		t.Logf("\t\tPASSED: %v", unit)
	} else {
		t.Fatalf("\t\tFAILED: %v (expected %v)", unit, unit_expected)
	}
}
