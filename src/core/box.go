package core

type Box struct {
	min, max Vec3
}

func NewBox(min, max Vec3) Box {
	return Box{min: min, max: max}
}
