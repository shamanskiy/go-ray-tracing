package vec3

type Vec3 struct {
	X, Y, Z float32
}

func (A Vec3) Add(B Vec3) Vec3 {
	return Vec3{A.X + B.X, A.Y + B.Y, A.Z + B.Z}
}

func (A Vec3) Sub(B Vec3) Vec3 {
	return Vec3{A.X - B.X, A.Y - B.Y, A.Z - B.Z}
}
