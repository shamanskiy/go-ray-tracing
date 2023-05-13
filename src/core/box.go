package core

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
	return NewBox(Vec3Min(box.min, other.min), Vec3Max(box.max, other.max))
}
