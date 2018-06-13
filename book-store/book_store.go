package bookstore

func Cost(basket []int) int {
    set, minCost := false, 0
    nextGroup := makeGroups(basket)
    for done, group := nextGroup(); !done; done, group = nextGroup() {
        if cost := cost(group); set == false || cost < minCost {
            set = true
            minCost = cost
        }
    }
    return minCost
}

func cost(books [][]int) (total int) {
    for _, group := range books {
        discount := 0
        switch len(group) {
        case 2:
            discount = 5
        case 3:
            discount = 10
        case 4:
            discount = 20
        case 5:
            discount = 25
        }
        total += len(group) * 800 * (100 - discount) / 100
    }
    return total
}

func makeGroups(items []int) func() (bool, [][]int) {
    return func() (bool, [][]int) {
        group := [][]int{}
        return true, group
    }
}
