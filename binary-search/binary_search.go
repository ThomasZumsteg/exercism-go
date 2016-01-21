package binarysearch

import "fmt"

/*SearchInts finds the index of an item in a list using a binary search.*/
func SearchInts(ints []int, find int) int {
	first, last := 0, len(ints)-1
	for first <= last {
		middle := (first + last) / 2
		if find <= ints[middle] {
			last = middle - 1
		} else {
			first = middle + 1
		}
	}
	return first
}

/*Message displays the results of a search in a nice way.*/
func Message(ints []int, find int) string {
	index := SearchInts(ints, find)

	atStart, atEnd := index == 0, index == len(ints)-1
	beyondEnd := index == len(ints)
	found := !beyondEnd && ints[index] == find

	switch {
	case 0 == len(ints):
		return "slice has no values"
	case found && atStart:
		return fmt.Sprintf("%d found at beginning of slice", find)
	case found && atEnd:
		return fmt.Sprintf("%d found at end of slice", find)
	case found:
		return fmt.Sprintf("%d found at index %d", find, index)
	case !found && atStart:
		return fmt.Sprintf("%d < all values", find)
	case !found && beyondEnd:
		return fmt.Sprintf("%d > all %d values", find, len(ints))
	case !found:
		return fmt.Sprintf("%d > %d at index %d, < %d at index %d",
			find, ints[index-1], index-1, ints[index], index)
	default:
		return ""
	}
}
