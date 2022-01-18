package glist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	intArr := []int{}
	for i := 0; i < 10; i++ {
		intArr = append(intArr, i)
	}

	intArr = Filter(intArr, func(item int, index int) bool {
		return item < 5
	})

	assert.Equal(len(intArr), 5)

}
