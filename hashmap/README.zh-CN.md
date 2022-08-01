# hashmap

[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/gog?label=go&logo=go&style=flat-square)](https://github.com/shalldie/gog)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/gog.svg)](https://pkg.go.dev/github.com/shalldie/gog/hashmap)
[![Build Status](https://img.shields.io/github/workflow/status/shalldie/gog/ci?label=test&logo=github&style=flat-square)](https://github.com/shalldie/gog/actions)
[![License](https://img.shields.io/github/license/shalldie/gog?logo=github&style=flat-square)](https://github.com/shalldie/gog)

Generic hashmap of Golang. Golang 泛型 hashmap 实现。

[English](./README.md) | [中文](./README.zh-CN.md)

Need v1.18+

<details><summary>示例</summary>
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

## 安装

```bash
go get github.com/shalldie/gog/hashmap
```

## 目录

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

生成 HashMap 的工厂方法.

### func Has

```go
func (hash *HashMap[K, V]) Has(key K) bool
```

`Has` 返回一个 bool 值，用来表明 HashMap 中是否存在指定元素。

### func Set

```go
func (hash *HashMap[K, V]) Set(key K, val V)
```

`Set` 方法为 HashMap 对象添加或更新一个指定了键（key）和值（value）的（新）键值对。

### func Get

```go
func (hash *HashMap[K, V]) Get(key K) V
```

`Get` 方法返回某个 HashMap 对象中的一个指定元素。

### func Delete

```go
func (hash *HashMap[K, V]) Delete(key K)
```

`Delete` 方法用于移除 HashMap 对象中指定的元素。

### func Keys

```go
func (hash *HashMap[K, V]) Keys() []K
```

`Keys` 返回一个新的切片（[]K），包含 HashMap 中所有的 key。

### func Values

```go
func (hash *HashMap[K, V]) Values() []V
```

`Values` 返回一个新的切片（[]V），包含 HashMap 中所有的 value。

### func Size

```go
func (hash *HashMap[K, V]) Size() int
```

`Size` 返回 HashMap 中键值对的数量。

### func ForEach

```go
func (hash *HashMap[K, V]) ForEach(action func(K, V))
```

`ForEach` 会对 HashMap 中每个键值对执行一次给定的函数。

### func Clear

```go
func (hash *HashMap[K, V]) Clear()
```

`Clear` 方法会移除 HashMap 对象中的所有元素。

### func Clone

```go
func (hash *HashMap[K, V]) Clone() *HashMap[K, V]
```

`Clone` 会创建一个 HashMap 的拷贝对象。

## License

MIT
