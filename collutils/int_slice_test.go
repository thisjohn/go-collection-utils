package collutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSlice_Filter(t *testing.T) {
	// Prepare
	data := []int{1, 2, 3}
	expectedData := []int{2}

	// Act
	newData := IntSlice(data).Filter(func(item int, i int) bool {
		return item%2 == 0
	})

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestIntSlice_Map(t *testing.T) {
	// Prepare
	data := []int{1, 2, 3}
	expectedData := []Any{2, 4, 6}

	// Act
	newData := IntSlice(data).Map(func(item int, i int) Any {
		return item * 2
	})

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestIntSlice_Reduce(t *testing.T) {
	// Prepare
	data := []int{1, 2, 3}
	expectedData := 6

	// Act
	newData := IntSlice(data).Reduce(func(acc Any, item int) Any {
		return acc.(int) + item
	}, 0)

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestIntSlice_Index(t *testing.T) {
	// Prepare
	data := []int{1, 2, 3}
	expectedIndex := 1

	// Act
	index := IntSlice(data).Index(2)

	// Assert
	assert.Equal(t, expectedIndex, index)
}

func TestIntSlice_Index_NotFound(t *testing.T) {
	// Prepare
	data := []int{1, 2, 3}
	expectedIndex := -1

	// Act
	index := IntSlice(data).Index(4)

	// Assert
	assert.Equal(t, expectedIndex, index)
}
