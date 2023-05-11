package core

import "github.com/chewxy/math32"

type Box struct {
	min, max Vec3
}

func NewBox(min, max Vec3) Box {
	return Box{min: min, max: max}
}

func NewInfiniteBox() Box {
	return Box{
		min: NewVec3(-Inf(), -Inf(), -Inf()),
		max: NewVec3(Inf(), Inf(), Inf()),
	}
}

func (box Box) Min() Vec3 {
	return box.min
}

func (box Box) Max() Vec3 {
	return box.max
}

func (box Box) Union(other Box) Box {
	newMin := NewVec3(
		math32.Min(box.min.X(), other.min.X()),
		math32.Min(box.min.Y(), other.min.Y()),
		math32.Min(box.min.Z(), other.min.Z()))
	newMax := NewVec3(
		math32.Max(box.max.X(), other.max.X()),
		math32.Max(box.max.Y(), other.max.Y()),
		math32.Max(box.max.Z(), other.max.Z()))
	return NewBox(newMin, newMax)
}
