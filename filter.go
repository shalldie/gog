package glist

func Filter[T any](list []T, f func(T, int) bool) []T {
	result := []T{}

	for index, item := range list {
		if f(item, index) {
			result = append(result, item)
		}
	}

	return result
}
