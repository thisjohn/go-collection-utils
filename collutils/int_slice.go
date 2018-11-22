package collutils

// IntSlice represents an list of ints
type IntSlice []int

// Filter returns an array of items that pass `fn` test
func (list IntSlice) Filter(fn func(item int, index int) bool) []int {
	var data []int
	for i, v := range list {
		if fn(v, i) {
			data = append(data, v)
		}
	}
	return data
}

// Map returns an array of new data by mapping each items
func (list IntSlice) Map(fn func(item int, index int) Any) []Any {
	var data []Any
	for i, v := range list {
		data = append(data, fn(v, i))
	}
	return data
}

// Reduce boils down items into a single value
func (list IntSlice) Reduce(fn func(acc Any, item int) Any, init Any) Any {
	acc := init
	for _, v := range list {
		acc = fn(acc, v)
	}
	return acc
}

// Index returns index of item in data
func (list IntSlice) Index(x int) int {
	for i, v := range list {
		if v == x {
			return i
		}
	}
	return -1
}
