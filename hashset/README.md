# hashset

[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/gog?label=go&logo=go&style=flat-square)](https://github.com/shalldie/gog)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/gog.svg)](https://pkg.go.dev/github.com/shalldie/gog)
[![Build Status](https://img.shields.io/github/workflow/status/shalldie/gog/ci?label=test&logo=github&style=flat-square)](https://github.com/shalldie/gog/actions)
[![License](https://img.shields.io/github/license/shalldie/gog?logo=github&style=flat-square)](https://github.com/shalldie/gog)

Generic hashset of Golang. Golang 泛型 hashset 实现。

Need v1.18+

## Example

```go
{
    hash := hashset.New[string, string]()
    hash.Set("name", "tom")
    hash.Set("name2", "tom")

	fmt.Println(hash.Get("name")) // "tom"
	fmt.Println(hash.Has("name")) // true
	fmt.Println(hash.Has("name2")) // false
}
```

## Installation

```bash
go get github.com/shalldie/gog/hashset
```

# Description

HashSet is a collection that contains no duplicate elements, the usage can refer to [HashMap](../hashmap).

HashSet 是一个不允许有重复元素的集合，使用方式可参考 [HashMap](../hashmap)。

## License

MIT
