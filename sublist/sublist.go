package sublist

import "reflect"

type Relation string

func Sublist(list1 []int, list2 []int) Relation {
	diff := len(list1) - len(list2)
	if diff == 0 && reflect.DeepEqual(list1, list2) {
		return "equal"
	} else if diff > 0 && Compare(list1, list2) {
		return "superlist"
	} else if diff < 0 && Compare(list2, list1) {
		return "sublist"
	}
	return "unequal"
}

func Compare(list1 []int, list2 []int) bool {
	for i, _ := range list1 {
		match := true
		for j, item := range list2 {
			if item != list1[i+j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
