package gs

import (
	"errors"
	"sort"
)

func Copy[T comparable](list []T) []T {
	result := make([]T, len(list))
	copy(result, list)
	return result
}

func Reverse[T comparable](list []T) []T {
	result := []T{}

	for i := len(list) - 1; i >= 0; i-- {
		result = append(result, list[i])
	}
	return result
}

func Sort[T comparable](list []T, less func(T, T) bool) []T {
	list = Copy(list)

	sort.Slice(list, func(i, j int) bool {
		return less(list[i], list[j])
	})
	return list
}

func Contains[T comparable](list []T, item T) bool {
	return IndexOf(list, item) > -1
}

func IndexOf[T comparable](list []T, item T) int {
	for index, cur := range list {
		if cur == item {
			return index
		}
	}

	return -1
}

func LastIndexOf[T comparable](list []T, item T) int {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == item {
			return i
		}
	}

	return -1
}

func Every[T comparable](list []T, predicate func(T, int) bool) bool {

	for index, item := range list {
		if !predicate(item, index) {
			return false
		}
	}

	return true
}

func Some[T comparable](list []T, predicate func(T, int) bool) bool {

	for index, item := range list {
		if predicate(item, index) {
			return true
		}
	}

	return false
}

func ForEach[T comparable](list []T, action func(T, int)) {

	for index, item := range list {
		action(item, index)
	}
}

func Map[T comparable, R comparable](list []T, f func(T, int) R) []R {
	result := []R{}

	for index, item := range list {
		result = append(result, f(item, index))
	}

	return result
}

func Filter[T comparable](list []T, predicate func(T, int) bool) []T {
	result := []T{}

	for index, item := range list {
		if predicate(item, index) {
			result = append(result, item)
		}
	}
	return result
}

func Reduce[T comparable, R comparable](list []T, f func(R, T, int) R, initialValue R) R {
	for index, item := range list {
		initialValue = f(initialValue, item, index)
	}

	return initialValue
}

func FindIndex[T comparable](list []T, predicate func(T, int) bool) int {
	for index, cur := range list {
		if predicate(cur, index) {
			return index
		}
	}

	return -1
}

func Find[T comparable](list []T, predicate func(T, int) bool) (item T, err error) {
	index := FindIndex(list, predicate)

	if index >= 0 {
		return list[index], nil
	}

	return item, errors.New("not found")
}
