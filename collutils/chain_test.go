package collutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	type Obj struct {
		item   string
		times2 int
	}

	data := []Any{"apple", "banana", "cherry"}

	res, err := NewChain(L(data)).
		Filter(func(item Any, index int) bool {
			return len(item.(string))%2 == 0
		}).
		Map(func(item Any, index int) Any {
			return &Obj{
				item:   item.(string),
				times2: len(item.(string)) * 2,
			}
		}).
		Reduce(func(acc Any, item Any) Any {
			return acc.(int) + item.(*Obj).times2
		}, 0).
		Value()

	assrt := assert.New(t)
	assrt.NoError(err)
	assrt.Equal(24, res)
}
