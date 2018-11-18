package collutils

type Any = interface{}
type StringSlice struct{}

func (ss StringSlice) Filter(data []string, fn func(string, int) bool) []string {
	var newData []string
	for i, v := range data {
		if fn(v, i) {
			newData = append(newData, v)
		}
	}
	return newData
}

func (ss StringSlice) Map(data []string, fn func(string, int) Any) []Any {
	var newData []Any
	for i, v := range data {
		newData = append(newData, fn(v, i))
	}
	return newData
}

func (ss StringSlice) Reduce(data []string, fn func(Any, string) Any, init Any) Any {
	res := init
	for _, v := range data {
		res = fn(res, v)
	}
	return res
}
