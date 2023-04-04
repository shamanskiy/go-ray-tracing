package filters_test

import (
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/src/core/slices/filters"
	"github.com/stretchr/testify/assert"
)

func TestGreaterOrEqualThan(t *testing.T) {
	filter := filters.GreaterOrEqualThan(1.)

	assert.True(t, filter(2.))
	assert.True(t, filter(1.))
	assert.False(t, filter(0.5))
}
