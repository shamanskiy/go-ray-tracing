package core

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestQuadEquation_TwoSolutions(t *testing.T) {
	t.Log("Given a quadratic equation 2x^2 - 6x + 4 = 0,")
	t.Log("\twe can solve it:")
	left, right, _ := SolveQuadraticEquation(2.0, -6.0, 4.0)

	utils.CheckResult(t, "left solution", left, float32(1.0))
	utils.CheckResult(t, "right solution", right, float32(2.0))
}

func TestQuadEquation_OneSolution(t *testing.T) {
	t.Log("Given a quadratic equation 2x^2 - 4x + 2 = 0,")
	t.Log("\twe can solve it:")
	left, right, _ := SolveQuadraticEquation(2.0, -4.0, 2.0)

	utils.CheckResult(t, "left solution", left, float32(1.0))
	utils.CheckResult(t, "right solution", right, float32(1.0))
}

func TestQuadEquation_NoSolutions(t *testing.T) {
	t.Log("Given a quadratic equation 2x^2 - 4x + 10 = 0,")
	t.Log("\tit should have no solutions:")
	left, right, err := SolveQuadraticEquation(2.0, -4.0, 10.0)
	if err != nil && left == 0.0 && right == 0.0 {
		t.Logf("\t\tPASSED: no solutions: %v", err)
	} else {
		t.Fatalf("\t\tFAILED: solutions are %v and %v, expected no solutions", left, right)
	}
}
