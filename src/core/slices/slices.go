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

func FindFirst[T constraints.Ordered](slice []T, filter func(T) bool) *T {
	for _, element := range slice {
		if filter(element) {
			return &element
		}
	}
	return nil
}
