package core

import "math/rand"

func RandomInUnitSphere() Vec3 {
	vec := Vec3{1.0, 0.0, 0.0}
	for vec.LenSqr() >= 1.0 {
		vec = Vec3{rand.Float32(), rand.Float32(), rand.Float32()}.Mul(2.0).Sub(Vec3{1.0, 1.0, 1.0})
	}
	return vec
}
