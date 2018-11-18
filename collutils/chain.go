package collutils

import (
	"fmt"
	"log"
	"reflect"
)

type job struct {
	methodName string
	args       []interface{}
}

type Chain struct {
	data []string
	jobs []job
}

func NewChain(data []string) *Chain {
	return &Chain{data: data}
}

func (c *Chain) Filter(fn func(string, int) bool) *Chain {
	return c.addJob("Filter", fn)
}

func (c *Chain) Map(fn func(string, int) Any) *Chain {
	return c.addJob("Map", fn)
}

func (c *Chain) Reduce(fn func(Any, string) Any, init Any) *Chain {
	return c.addJob("Reduce", fn, init)
}

func (c *Chain) addJob(methodName string, args ...interface{}) *Chain {
	c.jobs = append(c.jobs, job{
		methodName: methodName,
		args:       args,
	})
	return c
}

func (c *Chain) Value() (interface{}, error) {
	value := reflect.ValueOf(StringSlice(c.data))

	for _, job := range c.jobs {
		if value.Kind() != reflect.Slice {
			return nil, fmt.Errorf("chaining non-slice data")
		}

		var inputs = []reflect.Value{}
		for _, v := range job.args {
			inputs = append(inputs, reflect.ValueOf(v))
		}

		log.Println("Call", job.methodName)
		results := value.MethodByName(job.methodName).Call(inputs)
		result := results[0]

		if result.Kind() == reflect.Slice {
			var list []string
			var item string

			for i := 0; i < result.Len(); i++ {
				retAt := result.Index(i)
				kind := retAt.Kind()
				if kind == reflect.Interface {
					switch x := retAt.Interface().(type) {
					case string:
						kind = reflect.String
						item = x
					}
				} else if kind == reflect.String {
					item = retAt.String()
				}

				// TODO: Support more kinds
				switch kind {
				case reflect.String:
					list = append(list, item)
				default:
					return nil, fmt.Errorf("unhandled item kind %d", kind)
				}
			}
			value = reflect.ValueOf(StringSlice(list))
		} else {
			value = result
		}
	}

	return value.Interface(), nil
}
