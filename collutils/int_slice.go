package collutils

type IntSlice struct{}

func (IntSlice) Filter(data []int, fn func(int, int) bool) []int {
	var newData []int
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

func (IntSlice) Map(data []int, fn func(int, int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

func (IntSlice) Reduce(data []int, fn func(Any, int) Any, init Any) Any {
	res := init
	for _, v := range data {
		res = fn(res, v)
	}
	return res
}
