package binary

import (
	"errors"
	"fmt"
	"strconv"
)

/*ParseBinary converts a string of binary digits to the decial equivalant
or returns an error.*/
func ParseBinary(bin string) (int, error) {
	dec := 0
	for _, v := range bin {
		n, ok := strconv.Atoi(string(v))
		switch {
		case ok != nil:
			msg := fmt.Sprintf("\"%c\" is not a vaid digit", v)
			return 0, errors.New(msg)
		case n < 0 || 1 < n:
			msg := fmt.Sprintf("\"%d\" is not a binary digit", n)
			return 0, errors.New(msg)
		default:
			dec = dec*2 + n
		}
	}
	return dec, nil
}
