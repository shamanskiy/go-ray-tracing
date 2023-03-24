package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func TestRay_ShouldEvaluate(t *testing.T) {
	ray := core.NewRay(core.NewVec3(0.0, 0.0, 0.0), core.NewVec3(1.0, 2.0, 3.0))

	point := ray.Eval(2.0)

	assert.Equal(t, core.NewVec3(2.0, 4.0, 6.0), point)
}

func TestVec3_ShouldMultiplyElementwise(t *testing.T) {
	vecA, vecB := core.NewVec3(1., 2., 3.), core.NewVec3(4., 5., 6.)

	product := vecA.MulVec(vecB)

	assert.Equal(t, core.NewVec3(4., 10., 18), product)
}

func TestVec3_ShouldReflectAroundAxis(t *testing.T) {
	vector := core.NewVec3(3., -4, 0.)
	axis := core.NewVec3(0., 1., 0.)

	reflected := vector.Reflect(axis)

	assert.Equal(t, core.NewVec3(3., 4., 0.), reflected)
}

func TestVec3_ShouldDivideByScalar(t *testing.T) {
	A := core.NewVec3(1., 2., 3.)
	b := core.Real(2.)

	assert.Equal(t, core.NewVec3(0.5, 1., 1.5), A.Div(b))
}
