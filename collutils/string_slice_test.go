package collutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlice_Filter(t *testing.T) {
	// Prepare
	list := []string{"apple", "banana", "cherry"}
	expectedData := []string{"banana", "cherry"}

	// Act
	data := StringSlice(list).Filter(func(item string, i int) bool {
		return len(item)%2 == 0
	})

	// Assert
	assert.Equal(t, expectedData, data)
}

func TestStringSlice_Map(t *testing.T) {
	// Prepare
	list := []string{"apple", "banana", "cherry"}
	expectedData := []Any{5, 6, 6}

	// Act
	data := StringSlice(list).Map(func(item string, i int) Any {
		return len(item)
	})

	// Assert
	assert.Equal(t, expectedData, data)
}

func TestStringSlice_Reduce(t *testing.T) {
	// Prepare
	list := []string{"apple", "banana", "cherry"}
	expectedData := "@+apple+banana+cherry"

	// Act
	data := StringSlice(list).Reduce(func(acc Any, item string) Any {
		return acc.(string) + "+" + item
	}, "@")

	// Assert
	assert.Equal(t, expectedData, data)
}

func TestStringSlice_Compact(t *testing.T) {
	// Prepare
	list := []string{"apple", "", "banana", "cherry", "", ""}
	expectedData := []string{"apple", "banana", "cherry"}

	// Act
	data := StringSlice(list).Compact()

	// Assert
	assert.Equal(t, expectedData, data)
}

func TestStringSlice_Index(t *testing.T) {
	// Prepare
	list := []string{"apple", "banana", "cherry"}
	expectedIndex := 1

	// Act
	index := StringSlice(list).Index("banana")

	// Assert
	assert.Equal(t, expectedIndex, index)
}

func TestStringSlice_Index_NotFound(t *testing.T) {
	// Prepare
	list := []string{"apple", "banana", "cherry"}
	expectedIndex := -1

	// Act
	index := StringSlice(list).Index("dog")

	// Assert
	assert.Equal(t, expectedIndex, index)
}
