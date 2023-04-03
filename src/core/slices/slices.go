package slices

import "golang.org/x/exp/constraints"

func Filter[T any](slice []T, filter func(T) bool) []T {
	filteredSlice := make([]T, 0, len(slice))
	for _, element := range slice {
		if filter(element) {
			filteredSlice = append(filteredSlice, element)
		}
	}
	return filteredSlice
}

func FindFirstLargerOrEqualThan[T constraints.Ordered](slice []T, value T) *T {
	for _, element := range slice {
		if element >= value {
			return &element
		}
	}
	return nil
}

func GreaterOrEqualThan[T constraints.Ordered](value T) func(T) bool {
	return func(x T) bool {
		return x >= value
	}
}
