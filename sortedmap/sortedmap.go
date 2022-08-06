package sortedmap

import (
	"github.com/shalldie/gog/hashmap"
)

type SortedMap[K comparable, V comparable] struct {
	*hashmap.HashMap[K, V]
}

func New[K comparable, V comparable]() *SortedMap[K, V] {
	return &SortedMap[K, V]{
		HashMap: &hashmap.HashMap[K, V]{
			M:  make(map[K]V),
			SK: []K{},
		},
	}
}

func (sm *SortedMap[K, V]) Clone() *SortedMap[K, V] {
	newsm := New[K, V]()

	sm.ForEach(func(k K, v V) {
		newsm.Set(k, v)
	})

	return newsm
}
