package lsproduct

import (
	"fmt"
	"strconv"
)

//TestVersion is the unit tests that these functions are designed to pass
const TestVersion = 1

/*LargestSeriesProduct finds the largest product of a sequence of digits
of a given length, it also checks the lengths are possible*/
func LargestSeriesProduct(number string, span int) (int, *string) {
	if len(number) < span {
		err := fmt.Sprintf("Span (%d) needs to be smaller than the number (%s)", span, number)
		return 0, &err
	}
	largest := 0
	for start := 0; start <= len(number)-span; start++ {
		slice := 1
		for i := 0; i < span; i++ {
			n, _ := strconv.Atoi(string(number[i+start]))
			slice *= n
		}
		if largest <= slice {
			largest = slice
		}
	}
	return largest, nil
}
