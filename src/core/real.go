package core

import "github.com/chewxy/math32"

type Real = float32

func Inf() Real {
	return math32.Inf(1)
}

func Abs(v Real) Real {
	return math32.Abs(v)
}
