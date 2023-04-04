package filters

import "golang.org/x/exp/constraints"

func GreaterOrEqualThan[T constraints.Ordered](value T) func(T) bool {
	return func(x T) bool {
		return x >= value
	}
}
