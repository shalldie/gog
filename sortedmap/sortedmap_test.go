package sortedmap

import (
	"testing"

	"github.com/shalldie/gog/gs"
	"github.com/stretchr/testify/assert"
)

func TestSorted(t *testing.T) {
	assert := assert.New(t)

	sm := New[int, int]()

	for i := 0; i < 10; i++ {
		sm.Set(i, i)
	}

	index := 0
	sm.ForEach(func(key, val int) {
		assert.Equal(key, index)
		index += 1
	})

	gs.ForEach(sm.Keys(), func(key, index int) {
		assert.Equal(sm.Values()[index], sm.Get(key))
	})

}
