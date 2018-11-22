package collutils

// StringSlice represents an list of strings
type StringSlice []string

// Filter returns an array of items that pass `fn` test
func (list StringSlice) Filter(fn func(item string, index int) bool) []string {
	var data []string
	for i, v := range list {
		if fn(v, i) {
			data = append(data, v)
		}
	}
	return data
}

// Map returns an array of new data by mapping each items
func (list StringSlice) Map(fn func(item string, index int) Any) []Any {
	var data []Any
	for i, v := range list {
		data = append(data, fn(v, i))
	}
	return data
}

// Reduce boils down items into a single value
func (list StringSlice) Reduce(fn func(acc Any, item string) Any, init Any) Any {
	acc := init
	for _, v := range list {
		acc = fn(acc, v)
	}
	return acc
}

// Compact filters out empty string and returns an array of items
func (list StringSlice) Compact() []string {
	return list.Filter(func(item string, index int) bool {
		return item != ""
	})
}

// Index returns index of item in data
func (list StringSlice) Index(x string) int {
	for i, v := range list {
		if v == x {
			return i
		}
	}
	return -1
}
