package hashset

import (
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/gog/hashmap"
)

type HashSet[K comparable, V comparable] struct {
	hm *hashmap.HashMap[K, V]
}

func New[K comparable, V comparable]() *HashSet[K, V] {
	return &HashSet[K, V]{
		hm: hashmap.New[K, V](),
	}
}

func (hs *HashSet[K, V]) Has(key K) bool {
	return hs.hm.Has(key)
}

func (hs *HashSet[K, V]) Set(key K, val V) {
	if !gs.Contains(hs.Values(), val) {
		hs.hm.Set(key, val)
	}
}

func (hs *HashSet[K, V]) Get(key K) V {
	return hs.hm.Get(key)
}

func (hs *HashSet[K, V]) Delete(key K) {
	hs.hm.Delete(key)
}

func (hs *HashSet[K, V]) Keys() []K {
	return hs.hm.Keys()
}

func (hs *HashSet[K, V]) Values() []V {
	return hs.hm.Values()
}

func (hs *HashSet[K, V]) Size() int {
	return hs.hm.Size()
}

func (hs *HashSet[K, V]) ForEach(action func(K, V)) {
	hs.hm.ForEach(action)
}

func (hs *HashSet[K, V]) Clear() {
	hs.hm.Clear()
}

func (hs *HashSet[K, V]) Clone() *HashSet[K, V] {
	newhashset := New[K, V]()
	hs.ForEach(func(k K, v V) {
		newhashset.Set(k, v)
	})
	return newhashset
}
