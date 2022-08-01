package hashset

import (
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/gog/hashmap"
)

type HashSet[K comparable, V comparable] struct {
	hashmap.HashMap[K, V]
}

func New[K comparable, V comparable]() *HashSet[K, V] {
	return &HashSet[K, V]{
		*hashmap.New[K, V](),
	}
}

func (hs *HashSet[K, V]) Set(key K, val V) {
	if !gs.Contains(hs.Values(), val) {
		hs.M[key] = val
	}
}

func (hs *HashSet[K, V]) Clone() *HashSet[K, V] {
	newhs := New[K, V]()
	hs.ForEach(func(k K, v V) {
		newhs.Set(k, v)
	})
	return newhs
}
