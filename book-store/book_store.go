package bookstore

type bookGrouping [][]int

func Cost(basket []int) int {
	minCostList := []bookGrouping{New()}
	minCost, isSet := 0, false
	for _, book := range basket {
		newCostList := []bookGrouping{}
		minCost, isSet = 0, false
		for _, grouping := range minCostList {
			for location := 0; location <= len(grouping); location++ {
				if ok, newGrouping := grouping.Insert(book, location); ok {
					if !isSet || newGrouping.Cost() < minCost {
						newCostList = []bookGrouping{}
						minCost = newGrouping.Cost()
						isSet = true
					}
					if newGrouping.Cost() == minCost {
						newCostList = append(newCostList, *newGrouping)
					}
				}
			}
			minCostList = newCostList
		}
	}
	return minCost
}

func New() bookGrouping {
	return make(bookGrouping, 0)
}

func (books *bookGrouping) Copy() bookGrouping {
	newGroup := make(bookGrouping, len(*books))
	for g, group := range *books {
		newGroup[g] = make([]int, len(group))
		copy(newGroup[g], group)
	}
    return newGroup
}

func (books *bookGrouping) Insert(newBook, location int) (bool, *bookGrouping) {
    newGroup := books.Copy()
	if len(*books) == location &&
		(location == 0 || 0 < len((*books)[len(*books)-1])) {
		newGroup = append(newGroup, []int{newBook})
	} else if 0 <= location && location < len(*books) {
		for _, book := range newGroup[location] {
			if book == newBook {
				return false, nil
			}
		}
		newGroup[location] = append(newGroup[location], newBook)
	} else {
		return false, nil
	}
	return true, &newGroup
}

func (books *bookGrouping) Cost() (total int) {
	for _, bookGrouping := range *books {
		discount := 0
		switch len(bookGrouping) {
		case 2:
			discount = 5
		case 3:
			discount = 10
		case 4:
			discount = 20
		case 5:
			discount = 25
		}
		total += len(bookGrouping) * 800 * (100 - discount) / 100
	}
	return total
}
