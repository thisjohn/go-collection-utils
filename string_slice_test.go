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
	newData := StringSlice{}.Filter(data, func(s string, i int) bool {
		return len(s)%2 == 0
	})

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestStringSlice_Map(t *testing.T) {
	// Prepare
	data := []string{"apple", "banana", "cherry"}
	expectedData := []Any{5, 6, 6}

	// Act
	newData := StringSlice{}.Map(data, func(s string, i int) Any {
		return len(s)
	})

	// Assert
	assert.Equal(t, expectedData, newData)
}

func TestStringSlice_Reduce(t *testing.T) {
	// Prepare
	data := []string{"apple", "banana", "cherry"}
	expectedData := "@+apple+banana+cherry"

	// Act
	newData := StringSlice{}.Reduce(data, func(acc Any, s string) Any {
		return acc.(string) + "+" + s
	}, "@")

	// Assert
	assert.Equal(t, expectedData, newData)
}
