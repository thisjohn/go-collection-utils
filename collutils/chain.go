package collutils

import (
	"fmt"
	"reflect"
)

// TODO: Still under development

// Chainable is an interface for propagating collection methods
type Chainable interface {
	Push(item interface{}) Chainable

	Filter(fn func(item Any, index int) bool) []Any
	Map(fn func(item Any, index int) Any) []Any
	Reduce(fn func(acc Any, item Any) Any, init Any) Any
}

type job struct {
	methodName string
	args       []interface{}
}

// Chain wraps collection methods and calling them one by one until `Value` is called
type Chain struct {
	chainable Chainable
	jobs      []job
}

// NewChain creates a new `Chain`
func NewChain(chainable Chainable) *Chain {
	return &Chain{chainable: chainable}
}

// Filter adds job "Filter"
func (c *Chain) Filter(fn func(item Any, index int) bool) *Chain {
	return c.addJob("Filter", fn)
}

// Map adds job "Map"
func (c *Chain) Map(fn func(item Any, index int) Any) *Chain {
	return c.addJob("Map", fn)
}

// Reduce adds job "Reduce"
func (c *Chain) Reduce(fn func(acc Any, item Any) Any, init Any) *Chain {
	return c.addJob("Reduce", fn, init)
}

func (c *Chain) addJob(methodName string, args ...interface{}) *Chain {
	c.jobs = append(c.jobs, job{
		methodName: methodName,
		args:       args,
	})
	return c
}

// Value computes and returns chaining result
func (c *Chain) Value() (interface{}, error) {
	value := reflect.ValueOf(c.chainable)

	for _, job := range c.jobs {
		if value.Kind() != reflect.Slice {
			return nil, fmt.Errorf("chaining non-slice data")
		}

		var inputs = []reflect.Value{}
		for _, v := range job.args {
			inputs = append(inputs, reflect.ValueOf(v))
		}

		results := value.MethodByName(job.methodName).Call(inputs)
		result := results[0]

		if result.Kind() == reflect.Slice {
			chainable := func() Chainable {
				// NOTE: Alway return L now
				return L{}
			}()
			if chainable == nil {
				return nil, fmt.Errorf("unhandled kind from result: %s", job.methodName)
			}

			for i := 0; i < result.Len(); i++ {
				item := result.Index(i).Interface()
				chainable = chainable.Push(item)
			}

			c.chainable = chainable
			value = reflect.ValueOf(c.chainable)
		} else {
			value = result
		}
	}

	return value.Interface(), nil
}
