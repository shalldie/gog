package glist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert := assert.New(t)

	strArr := []string{"hello", "world"}

	lenArr := Map(strArr, func(item string, index int) int {
		return len(item)
	})

	for i := 0; i < len(strArr); i++ {
		assert.Equal(len(strArr[i]), lenArr[i])
	}

}
