package octal

import (
	"errors"
	"fmt"
)

/*ParseOctal converts a string of octal digits to a decimal number.*/
func ParseOctal(oct string) (int64, error) {
	var dec int64
	for i, octDigit := range oct {
		if octDigit < '0' || '7' < octDigit {
			err := fmt.Sprintf("Not valid character: %s[%c]%s",
				oct[:i], oct[i], oct[i+1:])
			return 0, errors.New(err)
		}
		dec *= 8
		dec += int64(octDigit - '0')
	}
	return dec, nil
}
