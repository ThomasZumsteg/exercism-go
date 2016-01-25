package lsproduct

import (
	"fmt"
	"strconv"
)

//TestVersion is the unit tests that these functions are designed to pass
const TestVersion = 2

/*LargestSeriesProduct finds the largest product of a sequence of digits
of a given length, it also checks the lengths are possible*/
func LargestSeriesProduct(number string, span int) (int, error) {
	switch {
	case len(number) < span:
		return 0, fmt.Errorf("Span (%d) needs to be smaller than the number (%s)", span, number)
	case span < 0:
		return 0, fmt.Errorf("Span needs to be positive: %d", span)
	}

	largest := 0
	for start := 0; start <= len(number)-span; start++ {
		slice := 1
		for i := 0; i < span; i++ {
			n, err := strconv.Atoi(string(number[i+start]))
			if err != nil {
				return 0, fmt.Errorf("%s(%c) does appear to be a valid digit", number[:i], number[i])
			}
			slice *= n
		}
		if largest <= slice {
			largest = slice
		}
	}
	return largest, nil
}
