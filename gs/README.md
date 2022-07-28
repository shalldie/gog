# gs

[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/gog?label=go&logo=go&style=flat-square)](https://github.com/shalldie/gog)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/gog.svg)](https://pkg.go.dev/github.com/shalldie/gog/gs)
[![Build Status](https://img.shields.io/github/workflow/status/shalldie/gog/ci?label=test&logo=github&style=flat-square)](https://github.com/shalldie/gog/actions)
[![License](https://img.shields.io/github/license/shalldie/gog?logo=github&style=flat-square)](https://github.com/shalldie/gog)

Generic functions for Slice of Golang. Golang 切片的常用泛型方法。

[English](./README.md) | [中文](./README.zh-CN.md)

Need v1.18+

## Installation

```bash
go get github.com/shalldie/gog/gs
```

## Index

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

## Functions

All functions are `immutable`.

### Copy

Copies elements from a source slice into a new slice.

```go
newlist := gs.Copy([]int{1})
```

### Reverse

Reverses the elements into a new slice.

```go
newlist := gs.Reverse([]int{1, 2, 3})
// [3 2 1]
```

### Sort

Sort a slice into a new slice.

```go
newlist = gs.Sort([]int{2, 5, 1, 3, 4, 5}, func(t1, t2 int) bool {
    return t1 < t2
})
// [1 2 3 4 5 5]
```

### IndexOf

Returns the index of the first occurrence of a value in a slice, or -1 if it is not present.

```go
list := []int{1, 2, 3, 2}

index2 := gs.IndexOf(list, 2)
// 1

index5 := gs.IndexOf(list, 5)
// -1
```

### LastIndexOf

Returns the index of the last occurrence of a specified value in a slice, or -1 if it is not present.

```go
list := []int{1, 2, 3, 2}

index2 := gs.LastIndexOf(list, 2)
// 3

index5 := gs.LastIndexOf(list, 5)
// -1
```

### Every

Determines whether all the members of a slice satisfy the specified test.

```go
allless4 := gs.Every([]int{1, 2, 3}, func(item, index int) bool {
    return item < 4
})
// true
```

### Some

Determines whether the specified callback function returns true for any element of a slice.

```go
hasEven := gs.Some([]int{5, 6, 7}, func(item, index int) bool {
    return item%2 == 0
})
// true
```

### ForEach

Performs the specified action for each element in a slice.

```go
gs.ForEach([]int{1, 2, 3}, func(item, index int) {
    println(item)
})
```

### Map

Calls a defined callback function on each element of a slice, and returns a slice that contains the results.

```go
list := gs.Map([]int{5, 6, 7}, func(item, index int) string {
    return strconv.Itoa(item)
})
// ["5" "6" "7"]
```

### Filter

Returns the elements of a slice that meet the condition specified in a callback function.

```go
list = gs.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(item, index int) bool {
    return item%2 == 0
})
// [2 4 6 8]
```

### Reduce

Calls the specified callback function for all the elements in a slice. The return value of the callback function is the accumulated result, and is provided as an argument in the next call to the callback function.

```go
sum := gs.Reduce([]int{1, 2, 3}, func(previousValue int, currentValue int, currentIndex int) int {
    return previousValue + currentValue
}, 0)
// 6
```

## FindIndex

Returns the index of the first element in the slice that satisfies the provided testing function. Otherwise, it returns -1

```go
index := gs.FindIndex([]int{1, 2, 3}, func(item, index int) bool {
    return item%2 == 0
})
// 1
```

## Find

Returns the value of the first element in the provided slice that satisfies the provided testing function.

```go
target, err := gs.Find([]int{1, 2, 3}, func(item, index int) bool {
    return item%2 == 0
})

if err != nil {
    println(target)
    // 2
}
```

## License

MIT
