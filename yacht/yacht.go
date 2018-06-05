package yacht

import (
	"reflect"
	"sort"
)

func Score(dice []int, catagory string) int {
	group := make(map[int]int)
	for _, die := range dice {
		if _, ok := group[die]; !ok {
			group[die] = 0
		}
		group[die]++
	}

	sort.Ints(dice)
	switch catagory {
	case "yacht":
		if num, ok := group[5]; ok && num == 5 {
			return 50
		}
	case "ones":
		if num, ok := group[1]; ok {
			return num
		}
	case "twos":
		if num, ok := group[2]; ok {
			return num * 2
		}
	case "threes":
		if num, ok := group[3]; ok {
			return num * 3
		}
	case "fours":
		if num, ok := group[4]; ok {
			return num * 4
		}
	case "fives":
		if num, ok := group[5]; ok {
			return num * 5
		}
	case "sixes":
		if num, ok := group[6]; ok {
			return num * 6
		}
	case "full house":
		count := []int{}
		score := 0
		for k := range group {
			count = append(count, group[k])
			score += k * group[k]
		}
		sort.Ints(count)
		if len(count) == 2 && count[0] == 2 && count[1] == 3 {
			return score
		}
	case "four of a kind":
		for k := range group {
			if group[k] >= 4 {
				return k * 4
			}
		}
	case "little straight":
		if reflect.DeepEqual(dice, []int{1, 2, 3, 4, 5}) {
			return 30
		}
	case "big straight":
		if reflect.DeepEqual(dice, []int{2, 3, 4, 5, 6}) {
			return 30
		}
	case "choice":
		sum := 0
		for _, die := range dice {
			sum += die
		}
		return sum
	}
	return 0
}
