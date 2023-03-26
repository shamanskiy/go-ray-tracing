package core_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core"
	"github.com/stretchr/testify/assert"
)

func TestQuadEquation_ShouldHaveTwoSolutions(t *testing.T) {
	solution := core.SolveQuadEquation(2.0, -6.0, 4.0)

	assert.EqualValues(t, 1, solution.Left)
	assert.EqualValues(t, 2, solution.Right)
	assert.False(t, solution.NoSolution)
}

func TestQuadEquation_ShouldHaveOneSolution(t *testing.T) {
	solution := core.SolveQuadEquation(2.0, -4.0, 2.0)

	assert.EqualValues(t, 1, solution.Left)
	assert.EqualValues(t, 1, solution.Right)
	assert.False(t, solution.NoSolution)
}

func TestQuadEquation_ShouldHaveNoSolutions(t *testing.T) {
	solution := core.SolveQuadEquation(2.0, -4.0, 10.0)

	assert.True(t, solution.NoSolution)
}
