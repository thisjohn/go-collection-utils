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
	newData := IntSlice{}.Filter(data, func(item int, i int) bool {
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
	newData := IntSlice{}.Map(data, func(item int, i int) Any {
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
	newData := IntSlice{}.Reduce(data, func(acc Any, item int) Any {
		return acc.(int) + item
	}, 0)

	// Assert
	assert.Equal(t, expectedData, newData)
}
