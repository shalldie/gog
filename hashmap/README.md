# hashmap

[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/gog.svg)](https://pkg.go.dev/github.com/shalldie/gog)
[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/gog?label=go&logo=go&style=flat-square)](https://github.com/shalldie/gog)
[![Build Status](https://img.shields.io/github/workflow/status/shalldie/gog/ci?label=test&logo=github&style=flat-square)](https://github.com/shalldie/gog/actions)
[![License](https://img.shields.io/github/license/shalldie/gog?logo=github&style=flat-square)](https://github.com/shalldie/gog)

Generic hashmap of Golang. Golang 泛型 hashmap 实现。

[English](./README.md) | [中文](./README.zh-CN.md)

Need v1.18+

<details><summary>Example</summary>
<p>

```go
{
    hash := hashmap.New[string, string]()
    hash.Set("name", "tom")

	fmt.Println(hash.Get("name")) // "tom"
	fmt.Println(hash.Has("name")) // true
}
```

</p>
</details>

## Installation

```bash
go get github.com/shalldie/gog/hashmap
```

## Index

- [type HashMap](#type-HashMap)
  - [func New\[K comparable, V comparable\]() \*HashMap\[K, V\]](#func-New)
  - [func (hash \*HashMap[K, V]) Has(key K) bool](#func-Has)
  - [func (hash \*HashMap[K, V]) Set(key K, val V)](#func-Set)
  - [func (hash \*HashMap[K, V]) Get(key K) V](#func-Get)
  - [func (hash \*HashMap[K, V]) Delete(key K)](#func-Delete)
  - [func (hash \*HashMap[K, V]) Keys() []K](#func-Keys)
  - [func (hash \*HashMap[K, V]) Values() []V](#func-Values)
  - [func (hash \*HashMap[K, V]) Size() int](#func-Size)
  - [func (hash \*HashMap[K, V]) ForEach(action func(K, V))](#func-ForEach)
  - [func (hash \*HashMap[K, V]) Clear()](#func-Clear)
  - [func (hash \*HashMap[K, V]) Clone() \*HashMap[K, V]](#func-Clone)

### type HashMap

```go
type HashMap[K comparable, V comparable] struct {
	// private fields
}
```

### func New

```go
func New[K comparable, V comparable]() *HashMap[K, V]
```

Factory to create a new HashMap.

### func Has

```go
func (hash *HashMap[K, V]) Has(key K) bool
```

`Has` returns a boolean indicating whether an element with the specified key exists or not.

### func Set

```go
func (hash *HashMap[K, V]) Set(key K, val V)
```

`Set` adds or updates an element with a specified key and a value to a HashMap object.

### func Get

```go
func (hash *HashMap[K, V]) Get(key K) V
```

`Get` returns a specified element from a HashMap object.

### func Delete

```go
func (hash *HashMap[K, V]) Delete(key K)
```

`Delete` removes the specified element from a HashMap object by key.

### func Keys

```go
func (hash *HashMap[K, V]) Keys() []K
```

`Keys` returns a slice of `K` that contains the keys for each element in the HashMap object.

### func Values

```go
func (hash *HashMap[K, V]) Values() []V
```

`Values` returns a slice of `V` that contains the values for each element in the Map object.

### func Size

```go
func (hash *HashMap[K, V]) Size() int
```

`Size` returns the number of elements in a HashMap object.

### func ForEach

```go
func (hash *HashMap[K, V]) ForEach(action func(K, V))
```

`ForEach` executes a provided function once per each key/value pair in the HashMap object.

### func Clear

```go
func (hash *HashMap[K, V]) Clear()
```

`Clear` removes all elements from a HashMap object.

### func Clone

```go
func (hash *HashMap[K, V]) Clone() *HashMap[K, V]
```

`Clone` creates a copy of the HashMap object.

## License

MIT
