package collutils

// L represents any value list
type L []Any

// Push item into data slice
func (list L) Push(item interface{}) Chainable {
	return append(list, item)
}

// Filter returns an array of items that pass `fn` test
func (list L) Filter(fn func(item Any, index int) bool) []Any {
	var data []Any
	for i, v := range list {
		if fn(v, i) {
			data = append(data, v)
		}
	}
	return data
}

// Map returns an array of new data by mapping each items
func (list L) Map(fn func(item Any, index int) Any) []Any {
	var data []Any
	for i, v := range list {
		data = append(data, fn(v, i))
	}
	return data
}

// Reduce boils down items into a single value
func (list L) Reduce(fn func(acc Any, item Any) Any, init Any) Any {
	acc := init
	for _, v := range list {
		acc = fn(acc, v)
	}
	return acc
}
