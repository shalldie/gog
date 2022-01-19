package glist

func Some[T any](list []T, f func(T, int) bool) bool {

	for index, item := range list {
		if f(item, index) {
			return true
		}
	}

	return false
}
