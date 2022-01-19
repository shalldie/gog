package glist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvery(t *testing.T) {
	assert := assert.New(t)

	strArr := []int{1, 2, 3, 4, 5, 6}

	less7 := Every(strArr, func(item int, _ int) bool {
		return item < 7
	})

	assert.True(less7)

}
