# gs

[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/gog?label=go&logo=go&style=flat-square)](https://github.com/shalldie/gog)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/gog.svg)](https://pkg.go.dev/github.com/shalldie/gog)
[![Build Status](https://img.shields.io/github/workflow/status/shalldie/gog/ci?label=test&logo=github&style=flat-square)](https://github.com/shalldie/gog/actions)
[![License](https://img.shields.io/github/license/shalldie/gog?logo=github&style=flat-square)](https://github.com/shalldie/gog)

Generic functions for Slice of Golang. Golang 切片的常用泛型方法。

[English](./README.md) | [中文](./README.zh-CN.md)

需要 v1.18+

## 安装

```bash
go get github.com/shalldie/gog/gs
```

## 目录

- [x] [Copy](#Copy)
- [x] [Reverse](#Reverse)
- [x] [Sort](#Sort)
- [x] [IndexOf](#IndexOf)
- [x] [LastIndexOf](#LastIndexOf)
- [x] [Every](#Every)
- [x] [Some](#Some)
- [x] [ForEach](#ForEach)
- [x] [Map](#Map)
- [x] [Filter](#Filter)
- [x] [Reduce](#Reduce)
- [x] [FindIndex](#FindIndex)
- [x] [Find](#Find)

## 方法

所有方法都是 `immutable` 的。

### Copy

从一个 `slice` 复制所有元素到一个新的 `slice`。

```go
newlist := gs.Copy([]int{1})
```

### Reverse

将 `slice` 中元素的位置颠倒，并返回新 `slice`。

```go
newlist := gs.Reverse([]int{1, 2, 3})
// [3 2 1]
```

### Sort

返回排序后的新 `slice`。

```go
newlist = gs.Sort([]int{2, 5, 1, 3, 4, 5}, func(t1, t2 int) bool {
    return t1 < t2
})
// [1 2 3 4 5 5]
```

### IndexOf

返回在 `slice` 中可以找到一个给定元素的第一个索引，如果不存在，则返回-1。

```go
list := []int{1, 2, 3, 2}

index2 := gs.IndexOf(list, 2)
// 1

index5 := gs.IndexOf(list, 5)
// -1
```

### LastIndexOf

返回指定元素在 `slice` 中的最后一个的索引，如果不存在则返回 -1。

```go
list := []int{1, 2, 3, 2}

index2 := gs.LastIndexOf(list, 2)
// 3

index5 := gs.LastIndexOf(list, 5)
// -1
```

### Every

测试一个 `slice` 内的所有元素是否都能通过某个指定函数的测试。

```go
allless4 := gs.Every([]int{1, 2, 3}, func(t, i int) bool {
    return t < 4
})
// true
```

### Some

测试 `slice` 中是不是至少有 1 个元素通过了被提供的函数测试。

```go
hasEven := gs.Some([]int{5, 6, 7}, func(t, i int) bool {
    return t%2 == 0
})
// true
```

### ForEach

对 `slice` 的每个元素执行一次给定的函数。

```go
gs.ForEach([]int{1, 2, 3}, func(t int, i int) {
    println(t)
})

```

### Map

创建一个新 `slice` ，其结果是该 `slice` 中的每个元素是调用一次提供的函数后的返回值。

```go
list := gs.Map([]int{5, 6, 7}, func(t int, i int) string {
    return strconv.Itoa(t)
})
// ["5" "6" "7"]
```

### Filter

创建一个新 `slice` , 其包含通过所提供函数实现的测试的所有元素。

```go
list = gs.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(t, i int) bool {
    return t%2 == 0
})
// [2 4 6 8]
```

### Reduce

对 `slice` 中的每个元素执行一个由您提供的 reducer 函数(升序执行)，将其结果汇总为单个返回值。

```go
sum := gs.Reduce([]int{1, 2, 3}, func(r int, t int, i int) int {
    return r + t
}, 0)
// 6
```

## FindIndex

返回 `slice` 中满足提供的测试函数的第一个元素的索引。若没有找到对应元素则返回-1。

```go
index := gs.FindIndex([]int{1, 2, 3}, func(t, i int) bool {
    return t%2 == 0
})
// 1
```

## Find

返回 `slice` 中满足提供的测试函数的第一个元素的值。

```go
item, err := gs.Find([]int{1, 2, 3}, func(t, i int) bool {
    return t%2 == 0
})

if err != nil {
    println(item)
    // 2
}
```

## License

MIT
