package vec3

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func (A Vec3) Add(B Vec3) Vec3 {
	return Vec3{A.X + B.X, A.Y + B.Y, A.Z + B.Z}
}

func (A Vec3) Sub(B Vec3) Vec3 {
	return Vec3{A.X - B.X, A.Y - B.Y, A.Z - B.Z}
}

func (A Vec3) MultVec(B Vec3) Vec3 {
	return Vec3{A.X * B.X, A.Y * B.Y, A.Z * B.Z}
}

func (A Vec3) Mult(b float64) Vec3 {
	return Vec3{A.X * b, A.Y * b, A.Z * b}
}

func (A Vec3) DivVec(B Vec3) Vec3 {
	return Vec3{A.X / B.X, A.Y / B.Y, A.Z / B.Z}
}

func (A Vec3) Div(b float64) Vec3 {
	return Vec3{A.X / b, A.Y / b, A.Z / b}
}

func (A Vec3) Dot(B Vec3) float64 {
	return A.X*B.X + A.Y*B.Y + A.Z*B.Z
}

func (A Vec3) Cross(B Vec3) Vec3 {
	return Vec3{A.Y*B.Z - A.Z*B.Y, -A.X*B.Z + A.Z*B.X, A.X*B.Y - A.Y*B.X}
}

func (A Vec3) LengthSquared() float64 {
	return A.X*A.X + A.Y*A.Y + A.Z*A.Z
}

func (A Vec3) Length() float64 {
	return math.Sqrt(A.LengthSquared())
}

func (A Vec3) Normalize() Vec3 {
	return A.Div(A.Length())
}

func (A Vec3) Elem(i int) float64 {
	switch i {
	case 0:
		return A.X
	case 1:
		return A.Y
	case 2:
		return A.Z
	default:
		return 0.0
	}
}
