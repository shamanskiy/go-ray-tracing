package random_test

import (
	"math"
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/src/core/random"
	"github.com/Shamanskiy/go-ray-tracer/test"
	"github.com/stretchr/testify/assert"
)

func TestRandomReal_ShouldBeBetweenZeroAndOne(t *testing.T) {
	randomGenerator := random.RandomGeneratedImpl{}
	numSamples := 1000000

	for i := 0; i < numSamples; i++ {
		randomReal := randomGenerator.Real()
		test.AssertInSemiInternal(t, randomReal, 0, 1)
	}
}

func TestRandomReal_ShouldBeApproximatelyUniform(t *testing.T) {
	randomGenerator := random.NewRandomGenerator()
	valueCounts := make(map[int]int)

	numSamples := 1000000
	numBins := 100
	maxDivergence := 0.10

	for i := 0; i < numSamples; i++ {
		randomReal := randomGenerator.Real()
		valueCounts[int(randomReal*float32(numBins))]++
	}

	for bin := 0; bin < numBins; bin++ {
		count := valueCounts[bin]
		divergence := math.Abs(float64(count-numSamples/numBins)) / float64(numSamples/numBins)
		assert.Less(t, divergence, maxDivergence, "bin %d, count %d, divergence %f", bin, count, divergence)
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
