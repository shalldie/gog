# hashmap

[![Go Version](https://img.shields.io/github/go-mod/go-version/shalldie/gog?label=go&logo=go&style=flat-square)](https://github.com/shalldie/gog)
[![Go Reference](https://pkg.go.dev/badge/github.com/shalldie/gog.svg)](https://pkg.go.dev/github.com/shalldie/gog/sortedmap)
[![Build Status](https://img.shields.io/github/workflow/status/shalldie/gog/ci?label=test&logo=github&style=flat-square)](https://github.com/shalldie/gog/actions)
[![License](https://img.shields.io/github/license/shalldie/gog?logo=github&style=flat-square)](https://github.com/shalldie/gog)

Generic sortedmap of Golang. Golang 泛型 sortedmap 实现。

Need v1.18+

<details><summary>Example</summary>
<p>

```go
{
    hash := sortedmap.New[int, int]()

    for i := 0; i < 10; i++ {
        sm.Set(i, i)
    }

    sm.ForEach(func(key, val int) {
        println(key, val)
    })

    // 0 1 2 3 4 5 6 7 8 9
}
```

</p>
</details>

## Installation

```bash
go get github.com/shalldie/gog/sortedmap
```

# Description

sortedmap is a collection that has sorted elements, the usage can refer to [HashMap](../hashmap).

sortedmap 是一个含有排序元素的集合，使用方式可参考 [HashMap](../hashmap)。

## License

MIT
