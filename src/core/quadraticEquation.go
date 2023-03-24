package core

import (
	"github.com/chewxy/math32"
)

type QuadEqSolution struct {
	Left       Real
	Right      Real
	NoSolution bool
}

func SolveQuadEquation(a, b, c Real) QuadEqSolution {
	d := b*b - 4*a*c
	if d >= 0 {
		dSqrt := math32.Sqrt(d)
		left := (-b - dSqrt) / (2 * a)
		right := (-b + dSqrt) / (2 * a)
		return QuadEqSolution{Left: left, Right: right, NoSolution: false}
	}

	return QuadEqSolution{NoSolution: true}
}
