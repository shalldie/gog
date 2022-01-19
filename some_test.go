package glist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	assert := assert.New(t)

	strArr := []int{1, 2, 3, 4, 5, 6}

	result := Some(strArr, func(item int, _ int) bool {
		return item > 5
	})

	assert.True(result)

}
