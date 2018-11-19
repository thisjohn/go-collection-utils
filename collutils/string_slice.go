package collutils

type StringSlice []string

func (data StringSlice) Filter(fn func(item string, index int) bool) []string {
	var newData []string
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

func (data StringSlice) Map(fn func(item string, index int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

func (data StringSlice) Reduce(fn func(acc Any, item string) Any, init Any) Any {
	res := init
	for _, v := range data {
		res = fn(res, v)
	}
	return res
}
