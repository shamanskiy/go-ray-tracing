package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/chewxy/math32"
	"github.com/stretchr/testify/assert"
)

func TestVec3_ShouldReturnCoordinates(t *testing.T) {
	vec := core.NewVec3(1., 2., 3.)

	assert.Equal(t, core.Real(1.), vec.X())
	assert.Equal(t, core.Real(2.), vec.Y())
	assert.Equal(t, core.Real(3.), vec.Z())
}

func TestVec3_ShouldAddVector(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(4., 5., 6.)

	sum := vecA.Add(vecB)

	assert.Equal(t, core.NewVec3(5., 7., 9.), sum)
}

func TestVec3_ShouldSubtractVector(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(4., 5., 6.)

	sum := vecA.Sub(vecB)

	assert.Equal(t, core.NewVec3(-3., -3., -3.), sum)
}

func TestVec3_ShouldMultiplyByScalar(t *testing.T) {
	vec := core.NewVec3(1., 2., 3.)
	b := core.Real(2.)

	product := vec.Mul(b)

	assert.Equal(t, core.NewVec3(2., 4., 6.), product)
}

func TestVec3_ShouldMultiplyByVectorElementwise(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(4., 5., 6.)

	product := vecA.MulVec(vecB)

	assert.Equal(t, core.NewVec3(4., 10., 18), product)
}

func TestVec3_ShouldComputeCrossProduct(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(4., 5., 6.)

	product := vecA.Cross(vecB)

	assert.Equal(t, core.NewVec3(-3., 6., -3.), product)
}

func TestVec3_ShouldDivideByScalar(t *testing.T) {
	vec := core.NewVec3(1., 2., 3.)
	b := core.Real(2.)

	quotient := vec.Div(b)

	assert.Equal(t, core.NewVec3(0.5, 1., 1.5), quotient)
}

func TestVec3_ShouldComputeDotProduct(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(4., 5., 6.)

	product := vecA.Dot(vecB)

	assert.Equal(t, core.Real(32), product)
}

func TestVec3_ShouldComputeLength(t *testing.T) {
	vec := core.NewVec3(1., 2., 3.)

	length := vec.Len()

	assert.InDelta(t, math32.Sqrt(14), length, core.Tolerance)
}

func TestVec3_ShouldComputeSquaredLength(t *testing.T) {
	vec := core.NewVec3(1., 2., 3.)

	length := vec.LenSqr()

	assert.EqualValues(t, 14, length)
}

func TestVec3_ShouldNormalize(t *testing.T) {
	vec := core.NewVec3(1., 2., 3.)

	normalized := vec.Normalize()

	assert.InDelta(t, 1., normalized.Len(), core.Tolerance)
}

func TestVec3_ShouldReflectAroundAxis(t *testing.T) {
	vector := core.NewVec3(3., -4, 0.)
	axis := core.NewVec3(0., 1., 0.)

	reflected := vector.Reflect(axis)

	assert.Equal(t, core.NewVec3(3., 4., 0.), reflected)
}

func TestVec3_ShouldSayTwoEqualVectorsAreWithinTolerance(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(1, 2, 3)

	assert.True(t, vecA.InDelta(vecB, core.Tolerance))
}

func TestVec3_ShouldSayTwoFarVectorsAreNotWithinTolerance(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(1+core.Tolerance, 2, 3)

	assert.False(t, vecA.InDelta(vecB, core.Tolerance))
}

func TestVec3_ShouldSayTwoCloseVectorsAreWithinTolerance(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(1+core.Tolerance/2, 2+core.Tolerance/2, 3+core.Tolerance/2)

	assert.True(t, vecA.InDelta(vecB, core.Tolerance))
}
