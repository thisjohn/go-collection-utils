package collutils

type AnySlice []Any

func (data AnySlice) Push(item interface{}) Chainable {
	return append(data, item)
}

func (data AnySlice) Filter(fn func(item Any, index int) bool) []Any {
	var newData []Any
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

func (data AnySlice) Map(fn func(item Any, index int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

func (data AnySlice) Reduce(fn func(acc Any, item Any) Any, init Any) Any {
	res := init
	for _, v := range data {
		res = fn(res, v)
	}
	return res
}
