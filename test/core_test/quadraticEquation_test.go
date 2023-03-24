package core

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/Shamanskiy/go-ray-tracer/test"
)

func TestQuadEquation_TwoSolutions(t *testing.T) {
	t.Log("Given a quadratic equation 2x^2 - 6x + 4 = 0,")
	t.Log("  we can solve it:")
	left, right, _ := core.SolveQuadraticEquation(2.0, -6.0, 4.0)

	test.CheckResult(t, "left solution", left, float32(1.0))
	test.CheckResult(t, "right solution", right, float32(2.0))
}

func TestQuadEquation_OneSolution(t *testing.T) {
	t.Log("Given a quadratic equation 2x^2 - 4x + 2 = 0,")
	t.Log("  we can solve it:")
	left, right, _ := core.SolveQuadraticEquation(2.0, -4.0, 2.0)

	test.CheckResult(t, "left solution", left, float32(1.0))
	test.CheckResult(t, "right solution", right, float32(1.0))
}

func TestQuadEquation_NoSolutions(t *testing.T) {
	t.Log("Given a quadratic equation 2x^2 - 4x + 10 = 0,")
	t.Log("  it should have no solutions:")
	left, right, err := core.SolveQuadraticEquation(2.0, -4.0, 10.0)

	t.Logf("\tError: %v\n", err)
	test.CheckResult(t, "left solution", left, core.Real(0.0))
	test.CheckResult(t, "right solution", right, core.Real(0.0))
}
