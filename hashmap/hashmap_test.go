package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetHas(t *testing.T) {
	assert := assert.New(t)

	hash := New[string, string]()
	hash.Set("name", "tom")

	assert.True(hash.Has("name"))
	assert.False(hash.Has("name2"))
}

func TestGetDelete(t *testing.T) {
	assert := assert.New(t)

	hash := New[string, string]()
	hash.Set("name", "tom")
	assert.Equal(hash.Get("name"), "tom")

	hash.Delete("name")
	assert.NotEqual(hash.Get("name"), "tom")
}

func TestSize(t *testing.T) {
	assert := assert.New(t)

	hash := New[string, string]()
	assert.Equal(hash.Size(), 0)

	hash.Set("name", "tom")
	assert.Equal(hash.Size(), 1)

	hash.Delete("name")
	assert.Equal(hash.Size(), 0)
}

func TestForEach(t *testing.T) {
	assert := assert.New(t)

	list := []string{"tom", "lily"}
	hash := New[int, string]()

	for index, item := range list {
		hash.Set(index, item)
	}

	hash.ForEach(func(k int, v string) {
		assert.Equal(hash.Get(k), list[k])
	})
	assert.Equal(hash.Size(), len(list))
}

func TestClear(t *testing.T) {
	assert := assert.New(t)

	list := []string{"tom", "lily"}
	hash := New[int, string]()
	assert.Equal(hash.Size(), 0)

	for index, item := range list {
		hash.Set(index, item)
	}

	assert.Equal(hash.Size(), len(list))

	hash.Clear()
	assert.Equal(hash.Size(), 0)
}

func TestClone(t *testing.T) {
	assert := assert.New(t)

	hash1 := New[string, string]()
	hash1.Set("name", "tom")

	hash2 := hash1.Clone()

	hash1.Delete("name")
	assert.NotEqual(hash1.Get("name"), "tom")
	assert.Equal(hash2.Get("name"), "tom")
}
