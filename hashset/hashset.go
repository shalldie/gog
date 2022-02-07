package hashset

import (
	"github.com/shalldie/gog/gs"
	"github.com/shalldie/gog/hashmap"
)

type HashSet[K comparable, V comparable] struct {
	hash *hashmap.HashMap[K, V]
}

func New[K comparable, V comparable]() *HashSet[K, V] {
	return &HashSet[K, V]{
		hash: hashmap.New[K, V](),
	}
}

func (hash *HashSet[K, V]) Has(key K) bool {
	return hash.hash.Has(key)
}

func (hash *HashSet[K, V]) Set(key K, val V) {
	if !gs.Contains(hash.Values(), val) {
		hash.hash.Set(key, val)
	}
}

func (hash *HashSet[K, V]) Get(key K) V {
	return hash.hash.Get(key)
}

func (hash *HashSet[K, V]) Delete(key K) {
	hash.hash.Delete(key)
}

func (hash *HashSet[K, V]) Keys() []K {
	return hash.hash.Keys()
}

func (hash *HashSet[K, V]) Values() []V {
	return hash.hash.Values()
}

func (hash *HashSet[K, V]) Size() int {
	return hash.hash.Size()
}

func (hash *HashSet[K, V]) ForEach(action func(K, V)) {
	hash.hash.ForEach(action)
	for _, key := range hash.Keys() {
		action(key, hash.Get(key))
	}
}

func (hash *HashSet[K, V]) Clear() {
	hash.hash.Clear()
}

func (hash *HashSet[K, V]) Clone() *HashSet[K, V] {
	newmap := New[K, V]()
	hash.ForEach(func(k K, v V) {
		newmap.Set(k, v)
	})
	return newmap
}
