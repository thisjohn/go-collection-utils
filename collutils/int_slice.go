package collutils

// IntSlice represents an list of ints
type IntSlice []int

// Filter returns an array of items that pass `fn` test
func (data IntSlice) Filter(fn func(item int, index int) bool) []int {
	var newData []int
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

// Map returns an array of new data by mapping each items
func (data IntSlice) Map(fn func(item int, index int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

// Reduce boils down items into a single value
func (data IntSlice) Reduce(fn func(acc Any, item int) Any, init Any) Any {
	acc := init
	for _, v := range data {
		acc = fn(acc, v)
	}
	return acc
}

// Index returns index of item in data
func (data IntSlice) Index(x int) int {
	for i, v := range data {
		if v == x {
			return i
		}
	}
	return -1
}
