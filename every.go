package glist

func Every[T any](list []T, f func(T, int) bool) bool {

	for index, item := range list {
		if !f(item, index) {
			return false
		}
	}

	return true
}
