package hashmap

import "github.com/shalldie/gog/gs"

type HashMap[K comparable, V comparable] struct {
	// internal map
	M map[K]V
	// sorted keys
	SK []K
}

func New[K comparable, V comparable]() *HashMap[K, V] {
	return &HashMap[K, V]{
		M: make(map[K]V),
	}
}

// 是否需要排序，通过初始化时候的 SK 来决定
func (hm *HashMap[K, V]) isSorted() bool {
	return hm.SK != nil
}

func (hm *HashMap[K, V]) Has(key K) bool {
	_, has := hm.M[key]
	return has
}

func (hm *HashMap[K, V]) Set(key K, val V) {
	// 如果需要排序，在 sorted keys 中添加该项
	if hm.isSorted() && !hm.Has(key) {
		hm.SK = append(hm.SK, key)
	}
	hm.M[key] = val
}

func (hm *HashMap[K, V]) Get(key K) V {
	return hm.M[key]
}

func (hm *HashMap[K, V]) Delete(key K) {
	// 如果需要排序，去掉 sorted keys 中的该项
	if hm.isSorted() {
		hm.SK = gs.Filter(hm.SK, func(item K, index int) bool {
			return item == key
		})
	}
	delete(hm.M, key)
}

func (hm *HashMap[K, V]) Keys() []K {
	keys := make([]K, hm.Size())
	index := 0
	hm.ForEach(func(k K, v V) {
		keys[index] = k
		index += 1
	})
	return keys
}

func (hm *HashMap[K, V]) Values() []V {
	vals := make([]V, hm.Size())
	index := 0
	hm.ForEach(func(k K, v V) {
		vals[index] = v
		index += 1
	})
	return vals
}

func (hm *HashMap[K, V]) Size() int {
	return len(hm.M)
}

func (hm *HashMap[K, V]) ForEach(action func(K, V)) {
	// 如果需要排序
	if hm.isSorted() {
		for _, key := range hm.SK {
			action(key, hm.Get(key))
		}
		return
	}
	// 乱序
	for key, val := range hm.M {
		action(key, val)
	}
}

func (hm *HashMap[K, V]) Clear() {
	hm.M = make(map[K]V)
	if hm.isSorted() {
		hm.SK = make([]K, 0)
	}
}

func (hm *HashMap[K, V]) Clone() *HashMap[K, V] {
	newhm := New[K, V]()

	// newhashmap.SK = hm.SK // in child struct

	hm.ForEach(func(k K, v V) {
		newhm.Set(k, v)
	})

	return newhm
}
