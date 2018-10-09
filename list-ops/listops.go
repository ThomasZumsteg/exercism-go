package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int
type IntList []int

func (list IntList) Foldr(fn binFunc, inital int) int {
	for i := len(list) - 1; i >= 0; i-- {
		inital = fn(list[i], inital)
	}
	return inital
}

func (list IntList) Foldl(fn binFunc, inital int) int {
	for _, item := range list {
		inital = fn(inital, item)
	}
	return inital
}

func (list IntList) Filter(fn predFunc) IntList {
	result := []int{}
	for _, item := range list {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func (list IntList) Length() int {
	return len(list)
}

func (list IntList) Map(fn unaryFunc) IntList {
	result := []int{}
	for _, item := range list {
		result = append(result, fn(item))
	}
	return result
}

func (list IntList) Reverse() IntList {
	result := []int{}
	for _, item := range list {
		result = append([]int{item}, result...)
	}
	return result
}

func (list IntList) Append(items IntList) IntList {
	return append(list, items...)
}

func (list IntList) Concat(lists []IntList) IntList {
	result := list
	for _, list := range lists {
		result = append(result, list...)
	}
	return result
}
