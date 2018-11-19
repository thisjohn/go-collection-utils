package collutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	data := []Any{"apple", "banana", "cherry"}

	res, err := NewChain(AnySlice(data)).
		Filter(func(item Any, index int) bool {
			return len(item.(string))%2 == 0
		}).
		Map(func(item Any, index int) Any {
			return len(item.(string)) * 2
		}).
		Reduce(func(acc Any, item Any) Any {
			return acc.(int) + item.(int)
		}, 0).
		Value()

	assrt := assert.New(t)
	assrt.NoError(err)
	assrt.Equal(24, res)
}
