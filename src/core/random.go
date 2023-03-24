package core

import (
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

type randomizer struct {
	on bool
}

var instance *randomizer

func Random() *randomizer {

	once.Do(func() {

		instance = &randomizer{on: true}
		rand.Seed(time.Now().Unix())
	})

	return instance
}

func (r *randomizer) VecInUnitSphere() Vec3 {
	if !r.on {
		return Vec3{0.0, 0.0, 0.0}
	}

	vec := Vec3{1.0, 0.0, 0.0}
	for vec.LenSqr() >= 1.0 {
		vec = Vec3{rand.Float32(), rand.Float32(), rand.Float32()}.Mul(2.0).Sub(Vec3{1.0, 1.0, 1.0})
	}
	return vec
}

func (r *randomizer) From01() Real {
	if !r.on {
		return 0.0
	}
	return rand.Float32()
}

func (r *randomizer) Vec3From01() Vec3 {
	if !r.on {
		return Vec3{0., 0., 0.}
	}
	return Vec3{rand.Float32(), rand.Float32(), rand.Float32()}
}

func (r *randomizer) Enable() {
	r.on = true
}

func (r *randomizer) Disable() {
	r.on = false
}
