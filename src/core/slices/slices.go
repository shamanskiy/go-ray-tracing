package slices

func Filter[T any](slice []T, filter func(T) bool) []T {
	filteredSlice := make([]T, 0, len(slice))
	for _, element := range slice {
		if filter(element) {
			filteredSlice = append(filteredSlice, element)
		}
	}
	return filteredSlice
}
