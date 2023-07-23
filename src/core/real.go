package core

import (
	"fmt"

	"github.com/chewxy/math32"
)

type Real = float32

func Inf() Real {
	return math32.Inf(1)
}

func Abs(v Real) Real {
	return math32.Abs(v)
}

func Sqrt(v Real) Real {
	return math32.Sqrt(v)
}

func Min(a, b Real) Real {
	return math32.Min(a, b)
}

type Interval struct {
	min, max Real
}

func NewInterval(min, max Real) Interval {
	if min > max+Tolerance {
		panic(fmt.Sprintf("new internal: min %f must be less or equal max %f", min, max))
	}
	return Interval{min: min, max: max}
}

func (i Interval) Contains(x Real) bool {
	return x >= i.min && x <= i.max
}

func (i Interval) ContainsStrictly(x Real) bool {
	return x > i.min && x < i.max
}
