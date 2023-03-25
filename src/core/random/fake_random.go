package random

import "github.com/Shamanskiy/go-ray-tracer/src/core"

type FakeRandomGenerator struct {
}

func NewFakeRandomGenerator() FakeRandomGenerator {
	return FakeRandomGenerator{}
}

func (f FakeRandomGenerator) Real() core.Real {
	return 0
}

func (f FakeRandomGenerator) Vec3() core.Vec3 {
	return core.NewVec3(0, 0, 0)
}

func (f FakeRandomGenerator) Vec3InUnitSphere() core.Vec3 {
	return core.NewVec3(0, 0, 0)
}