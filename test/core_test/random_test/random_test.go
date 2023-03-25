package random_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

func TestRandomReal(t *testing.T) {
	randomGenerator := random.RandomGeneratedImpl{}

	for i := 0; i < 10; i++ {
		randomReal := randomGenerator.Real()
		test.AssertInSemiInternal(t, randomReal, 0, 1)
	}
}

func TestRandomVec3(t *testing.T) {
	randomGenerator := random.RandomGeneratedImpl{}

	for i := 0; i < 10; i++ {
		randomVec3 := randomGenerator.Vec3()
		test.AssertInSemiInternal(t, randomVec3.X(), 0, 1)
		test.AssertInSemiInternal(t, randomVec3.Y(), 0, 1)
		test.AssertInSemiInternal(t, randomVec3.Z(), 0, 1)
	}
}

func TestRandomVec3InUnitSphere(t *testing.T) {
	randomGenerator := random.RandomGeneratedImpl{}

	for i := 0; i < 10; i++ {
		randomSphereVec := randomGenerator.Vec3InUnitSphere()
		assert.Less(t, randomSphereVec.LenSqr(), core.Real(1))
	}
}