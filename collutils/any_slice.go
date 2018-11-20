package collutils

// AnySlice represents an list of items
type AnySlice []Any

// Push item into data slice
func (data AnySlice) Push(item interface{}) Chainable {
	return append(data, item)
}

// Filter returns an array of items that pass `fn` test
func (data AnySlice) Filter(fn func(item Any, index int) bool) []Any {
	var newData []Any
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

// Map returns an array of new data by mapping each items
func (data AnySlice) Map(fn func(item Any, index int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

// Reduce boils down items into a single value
func (data AnySlice) Reduce(fn func(acc Any, item Any) Any, init Any) Any {
	acc := init
	for _, v := range data {
		acc = fn(acc, v)
	}
	return acc
}
