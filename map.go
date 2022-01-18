package glist

func Map[T any, R any](list []T, f func(T, int) R) []R {
	result := []R{}

	for index, item := range list {
		result = append(result, f(item, index))
	}

	return result
}
