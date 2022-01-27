package gs

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	assert := assert.New(t)

	arr := []int{1}
	list := Copy(arr)

	arr[0] = 2
	assert.NotEqual(arr[0], list[0])
	assert.Equal(len(list), 1)
}

func TestReverse(t *testing.T) {
	assert := assert.New(t)

	arr := []int{1, 2, 3}
	list := Reverse(arr)

	for i := 1; i < len(list); i++ {
		assert.Equal(arr[i], list[len(arr)-i-1])
	}
}

func TestSort(t *testing.T) {
	assert := assert.New(t)

	list := []int{2, 5, 1, 3, 4, 5}

	list = Sort(list, func(t1, t2 int) bool {
		return t1 < t2
	})

	for i := 1; i < len(list); i++ {
		assert.True(list[i] >= list[i-1])
	}
}

func TestIndexOf_LastIndexOf(t *testing.T) {
	assert := assert.New(t)

	list := []int{1, 2, 3, 2}

	assert.Equal(IndexOf(list, 2), 1)
	assert.Equal(IndexOf(list, 5), -1)

	assert.Equal(LastIndexOf(list, 2), 3)
	assert.Equal(LastIndexOf(list, 5), -1)
}

func TestEvery(t *testing.T) {
	assert := assert.New(t)

	list := []int{1, 2, 3}

	assert.True(
		Every(list, func(t, i int) bool {
			return t < 4
		}),
	)

	assert.False(
		Every(list, func(t, i int) bool {
			return t < 3
		}),
	)
}

func TestSome(t *testing.T) {
	assert := assert.New(t)

	list := []int{1, 2, 3}

	assert.True(
		Some(list, func(t, i int) bool {
			return t == 2
		}),
	)

	assert.False(
		Some(list, func(t, i int) bool {
			return t == 5
		}),
	)
}

func TestForEach(t *testing.T) {
	assert := assert.New(t)

	list := []int{1, 2, 3}
	list2 := []int{}

	ForEach(list, func(t int, i int) {
		list2 = append(list2, t)
	})

	assert.Equal(len(list), len(list))
}

func TestMap(t *testing.T) {
	assert := assert.New(t)

	list := []int{5, 6, 7}

	list2 := Map(list, func(t int, i int) string {
		return strconv.Itoa(t)
	})

	assert.Equal(
		len(
			strings.Join(list2, ""),
		),
		3,
	)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	list := []int{1, 2, 3, 4, 5, 6, 7, 8}

	list = Filter(list, func(t, i int) bool {
		return t%2 == 0
	})

	ForEach(list, func(t, i int) {
		assert.Equal(t%2, 0)
	})
}

func TestReduce(t *testing.T) {
	assert := assert.New(t)

	sum100arr := []int{}
	sum100 := 0

	for i := 1; i <= 100; i++ {
		sum100 += i
		sum100arr = append(sum100arr, i)
	}

	sumReduce := Reduce(sum100arr, func(r int, t int, i int) int {
		return r + t
	}, 0)

	assert.Equal(sum100, sumReduce)
}

func TestFindIndex(t *testing.T) {
	assert := assert.New(t)

	list := []int{1, 2, 3}

	assert.Equal(
		FindIndex(list, func(t, i int) bool {
			return t%2 == 0
		}),
		1,
	)

	assert.Equal(
		FindIndex(list, func(t, i int) bool {
			return t == 5
		}),
		-1,
	)
}

func TestFind(t *testing.T) {
	assert := assert.New(t)

	list := []int{5, 6, 7}

	item, err := Find(list, func(t, i int) bool {
		return t%2 == 0
	})

	assert.Nil(err)
	assert.Equal(item, 6)

	item, err = Find(list, func(t, i int) bool {
		return t == 8
	})

	assert.NotNil(err)
}
