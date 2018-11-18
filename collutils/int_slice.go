package collutils

type IntSlice []int

func (data IntSlice) Filter(fn func(int, int) bool) []int {
	var newData []int
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

func (data IntSlice) Map(fn func(int, int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

func (data IntSlice) Reduce(fn func(Any, int) Any, init Any) Any {
	res := init
	for _, v := range data {
		res = fn(res, v)
	}
	return res
}
