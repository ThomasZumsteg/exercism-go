package summultiples

/*MultipleSummer create a function that calculates the sum
of all unique multiples of numbers less than some value */
func MultipleSummer(multiples ...int) func(int) int {
	/*A function that calculates the sum of all unique multiples of multiples*/
	return func(limit int) int {
		total := 0
		for n := 1; n < limit; n++ {
			isMultiple := func(d int) bool { return n%d == 0 }
			if any(multiples, isMultiple) {
				total += n
			}
		}
		return total
	}
}

/*any tests if any items in a list cause the test to pass*/
func any(list []int, test func(int) bool) bool {
	for _, v := range list {
		if test(v) {
			return true
		}
	}
	return false
}
