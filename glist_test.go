package glist

import (
	"glist/each"
	"testing"
)

func TestEach(t *testing.T) {

	// each.Each([]int{1, 2, 3}, func(i int, item int) {
	// 	print(item)
	// })

	each.Each([]string{"hello", "world"}, func(item string) {
		println(item)
	})
}
