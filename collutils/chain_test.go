package collutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	data := []string{"apple", "banana", "cherry"}

	res, err := NewChain(data).
		Filter(func(s string, i int) bool {
			return len(s)%2 == 0
		}).
		Map(func(s string, i int) Any {
			return s[1:]
		}).
		Reduce(func(acc Any, s string) Any {
			return acc.(int) + len(s)
		}, 0).
		Value()

	assrt := assert.New(t)
	assrt.NoError(err)
	assrt.Equal(10, res)
}
