package collutils

type IntSlice []int

func (data IntSlice) Filter(fn func(item int, index int) bool) []int {
	var newData []int
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

func (data IntSlice) Map(fn func(item int, index int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

func (data IntSlice) Reduce(fn func(acc Any, item int) Any, init Any) Any {
	res := init
	for _, v := range data {
		res = fn(res, v)
	}
	return res
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
