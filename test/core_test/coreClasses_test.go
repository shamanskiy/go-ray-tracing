package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

func TestRay_ShouldEvaluate(t *testing.T) {
	ray := core.Ray{
		Origin:    core.Vec3{0.0, 0.0, 0.0},
		Direction: core.Vec3{1.0, 2.0, 3.0},
	}

	point := ray.Eval(2.0)

	assert.Equal(t, core.Vec3{2.0, 4.0, 6.0}, point)
}

func TestVec3_ShouldMultiplyElementwise(t *testing.T) {
	vecA, vecB := core.Vec3{1., 2., 3.}, core.Vec3{4., 5., 6.}

	product := core.MulElem(vecA, vecB)

	assert.Equal(t, core.Vec3{4., 10., 18}, product)
}

func TestColor_ShouldMultiplyElementWise(t *testing.T) {
	colorA, colorB := core.White, core.Red

	product := core.MulElem(colorA, colorB)

	assert.Equal(t, core.Red, product)
}

func TestVec3_ShouldReflectAroundAxis(t *testing.T) {
	vector := core.Vec3{3., -4, 0.}
	axis := core.Vec3{0., 1., 0.}

	reflected := core.Reflect(vector, axis)

	assert.Equal(t, core.Vec3{3., 4., 0.}, reflected)
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

func TestDiv(t *testing.T) {
	A := core.Vec3{1., 2., 3.}
	b := core.Real(2.)
	t.Logf("Given a vector %v and a scalar %v,\n", A, b)
	t.Log("  we can divide the vector by the scalar:")
	test.CheckResult(t, "Division result", core.Div(A, b), core.Vec3{0.5, 1., 1.5})
}
