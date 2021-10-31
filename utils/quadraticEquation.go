package utils

import (
	"fmt"

	"github.com/chewxy/math32"
)

func SolveQuadraticEquation(a, b, c float32) (left float32, right float32, err error) {
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
