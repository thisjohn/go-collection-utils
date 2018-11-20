package collutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlice_Filter(t *testing.T) {
	// Prepare
	data := []string{"apple", "banana", "cherry"}
	expectedData := []string{"banana", "cherry"}

	// Act
	newData := StringSlice(data).Filter(func(item string, i int) bool {
		return len(item)%2 == 0
	})

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestStringSlice_Map(t *testing.T) {
	// Prepare
	data := []string{"apple", "banana", "cherry"}
	expectedData := []Any{5, 6, 6}

	// Act
	newData := StringSlice(data).Map(func(item string, i int) Any {
		return len(item)
	})

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestStringSlice_Reduce(t *testing.T) {
	// Prepare
	data := []string{"apple", "banana", "cherry"}
	expectedData := "@+apple+banana+cherry"

	// Act
	newData := StringSlice(data).Reduce(func(acc Any, item string) Any {
		return acc.(string) + "+" + item
	}, "@")

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestStringSlice_Compact(t *testing.T) {
	// Prepare
	data := []string{"apple", "", "banana", "cherry", "", ""}
	expectedData := []string{"apple", "banana", "cherry"}

	// Act
	newData := StringSlice(data).Compact()

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestStringSlice_Index(t *testing.T) {
	// Prepare
	data := []string{"apple", "banana", "cherry"}
	expectedIndex := 1

	// Act
	index := StringSlice(data).Index("banana")

	// Assert
	assert.Equal(t, expectedIndex, index)
}

func TestStringSlice_Index_NotFound(t *testing.T) {
	// Prepare
	data := []string{"apple", "banana", "cherry"}
	expectedIndex := -1

	// Act
	index := StringSlice(data).Index("dog")

	// Assert
	assert.Equal(t, expectedIndex, index)
}
