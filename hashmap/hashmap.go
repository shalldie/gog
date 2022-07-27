package hashmap

type HashMap[K comparable, V comparable] struct {
	m map[K]V
}

func New[K comparable, V comparable]() *HashMap[K, V] {
	return &HashMap[K, V]{
		m: make(map[K]V),
	}
}

func (hm *HashMap[K, V]) Has(key K) bool {
	_, has := hm.m[key]
	return has
}

func (hm *HashMap[K, V]) Set(key K, val V) {
	hm.m[key] = val
}

func (hm *HashMap[K, V]) Get(key K) V {
	return hm.m[key]
}

func (hm *HashMap[K, V]) Delete(key K) {
	delete(hm.m, key)
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
	return len(hm.m)
}

func (hm *HashMap[K, V]) ForEach(action func(K, V)) {
	for key, val := range hm.m {
		action(key, val)
	}
}

func (hm *HashMap[K, V]) Clear() {
	hm.m = make(map[K]V)
}

func (hm *HashMap[K, V]) Clone() *HashMap[K, V] {
	newhashmap := New[K, V]()
	hm.ForEach(func(k K, v V) {
		newhashmap.Set(k, v)
	})
	return newhashmap
}
