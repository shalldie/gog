package hashmap

type HashMap[K comparable, V comparable] struct {
	m map[K]V
}

func New[K comparable, V comparable]() *HashMap[K, V] {
	return &HashMap[K, V]{
		m: make(map[K]V),
	}
}

func (hash *HashMap[K, V]) Has(key K) bool {
	_, has := hash.m[key]
	return has
}

func (hash *HashMap[K, V]) Set(key K, val V) {
	hash.m[key] = val
}

func (hash *HashMap[K, V]) Get(key K) V {
	return hash.m[key]
}

func (hash *HashMap[K, V]) Delete(key K) {
	delete(hash.m, key)
}

func (hash *HashMap[K, V]) Keys() []K {
	keys := make([]K, hash.Size())
	index := 0
	hash.ForEach(func(k K, v V) {
		keys[index] = k
		index += 1
	})
	return keys
}

func (hash *HashMap[K, V]) Values() []V {
	vals := make([]V, hash.Size())
	index := 0
	hash.ForEach(func(k K, v V) {
		vals[index] = v
		index += 1
	})
	return vals
}

func (hash *HashMap[K, V]) Size() int {
	return len(hash.m)
}

func (hash *HashMap[K, V]) ForEach(action func(K, V)) {
	for key, val := range hash.m {
		action(key, val)
	}
}

func (hash *HashMap[K, V]) Clear() {
	hash.m = make(map[K]V)
}

func (hash *HashMap[K, V]) Clone() *HashMap[K, V] {
	newhash := New[K, V]()
	hash.ForEach(func(k K, v V) {
		newhash.Set(k, v)
	})
	return newhash
}
