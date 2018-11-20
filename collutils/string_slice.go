package collutils

// StringSlice represents an list of strings
type StringSlice []string

// Filter returns an array of items that pass `fn` test
func (data StringSlice) Filter(fn func(item string, index int) bool) []string {
	var newData []string
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

// Map returns an array of new data by mapping each items
func (data StringSlice) Map(fn func(item string, index int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

// Reduce boils down items into a single value
func (data StringSlice) Reduce(fn func(acc Any, item string) Any, init Any) Any {
	acc := init
	for _, v := range data {
		acc = fn(acc, v)
	}
	return acc
}

// Compact filters out empty string and returns an array of items
func (data StringSlice) Compact() []string {
	return data.Filter(func(item string, index int) bool {
		return item != ""
	})
}

// Index returns index of item in data
func (data StringSlice) Index(x string) int {
	for i, v := range data {
		if v == x {
			return i
		}
	}
	return -1
}
