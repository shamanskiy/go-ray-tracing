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
