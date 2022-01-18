package glist

func Each[T any](list []T, f func(T, int)) {
	for index, item := range list {
		f(item, index)
	}
}
