package core

import (
	"fmt"

	"github.com/chewxy/math32"
)

func SolveQuadraticEquation(a, b, c Real) (left Real, right Real, err error) {
	d := b*b - 4*a*c
	if d >= 0 {
		dSqrt := math32.Sqrt(d)
		left = (-b - dSqrt) / (2 * a)
		right = (-b + dSqrt) / (2 * a)
		err = nil
		return
	}

	err = fmt.Errorf("negative discriminant: %v", d)
	return
}
