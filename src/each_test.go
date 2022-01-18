package glist

import (
	"testing"
)

func TestEach(t *testing.T) {

	Each([]string{"hello", "world"}, func(item string, index int) {

	})

	t.Log("done")

}
