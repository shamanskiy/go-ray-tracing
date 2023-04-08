package random

import (
	"hash/maphash"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
)

type RandomGenerator interface {
	Real() core.Real
	Vec3() core.Vec3
	Vec3InUnitSphere() core.Vec3
}

type RandomGeneratedImpl struct{}

func NewRandomGenerator() RandomGenerator {
	return RandomGeneratedImpl{}
}

// https://qqq.ninja/blog/post/fast-threadsafe-randomness-in-go/
func (r RandomGeneratedImpl) Real() core.Real {
	outUint64 := new(maphash.Hash).Sum64()
	outFloat64 := float64(outUint64) / float64(1<<64)
	if outFloat64 >= 1 {
		outFloat64 = 0.
	}
	return float32(outFloat64)
}

func (r RandomGeneratedImpl) Vec3() core.Vec3 {
	return core.NewVec3(r.Real(), r.Real(), r.Real())
}

func (r RandomGeneratedImpl) Vec3InUnitSphere() core.Vec3 {
	unitDiagVec := core.NewVec3(1, 1, 1)

	vec := core.NewVec3(1, 0, 0)
	for vec.LenSqr() >= 1 {
		vec = r.Vec3().Mul(2).Sub(unitDiagVec)
	}
	return vec
}
