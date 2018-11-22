package collutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetString(t *testing.T) {
	data := M{
		"A": "apple",
	}
	expectedStr := "apple"

	str := data.GetStringAndTrim("A")

	assert.Equal(t, expectedStr, str)
}

func TestGetStringAndTrim(t *testing.T) {
	data := M{
		"A": " apple   ",
	}
	expectedStr := "apple"

	str := data.GetStringAndTrim("A")

	assert.Equal(t, expectedStr, str)
}

func TestGetStringNoKey(t *testing.T) {
	data := M{
		"A": "apple",
	}
	expectedStr := ""

	str := data.GetStringAndTrim("B")

	assert.Equal(t, expectedStr, str)
}

func TestPick(t *testing.T) {
	data := M{
		"A":   "apple",
		"B":   "banana",
		"C":   "cherry",
		"one": 1,
	}
	expectedList := L{"apple", "cherry", "banana", 1}

	list := data.Pick("A", "C", "B", "one")

	assert.Equal(t, expectedList, list)
}

func TestExtend(t *testing.T) {
	data := M{"A": 1}
	expectedData := M{"A": 1, "B": 2}

	data.Extend(M{"B": 2})

	assert.Equal(t, expectedData, data)
}

func TestPluck(t *testing.T) {
	data := []M{
		{"A": 1, "B": 2},
		{"A": 10, "B": 20},
		{"A": 100, "B": 200},
	}
	expectedList := L{1, 10, 100}

	list := Pluck(data, "A")

	assert.Equal(t, expectedList, list)
}
