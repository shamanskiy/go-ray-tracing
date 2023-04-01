package slices_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core/slices"
	"github.com/stretchr/testify/assert"
)

func TestSlices_ShouldFilterSlice(t *testing.T) {
	slice := []int{-2, -1, 0, 1, 2}
	filter := func(elem int) bool {
		return elem >= 0
	}

	filteredSlice := slices.Filter(slice, filter)

	assert.Equal(t, []int{0, 1, 2}, filteredSlice)
}
