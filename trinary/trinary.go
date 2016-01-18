package trinary

import "fmt"

/*ParseTrinary converts a trinary string to an integer.*/
func ParseTrinary(tri string) (int64, error) {
	var total int64
	for i, t := range tri {
		total *= 3
		switch t {
		case '0':
			total += 0
		case '1':
			total += 1
		case '2':
			total += 2
		default:
			return 0, fmt.Errorf("Not a trinary digit: %q[%q]", tri[:i], tri[i:])
		}
		if total < 0 {
			return 0, fmt.Errorf("Overflow error, number to large")
		}
	}
	return total, nil
}
