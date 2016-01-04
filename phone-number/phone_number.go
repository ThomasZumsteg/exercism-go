package phonenumber

import (
	"errors"
	"fmt"
)

/*Number extracts the digits of a phone number from surrounding text.*/
func Number(number string) (string, error) {
	digits := ""
	for _, c := range number {
		if '0' <= c && c <= '9' {
			digits += string(c)
		}
	}
	switch {
	case len(digits) == 10:
		return digits, nil
	case len(digits) == 11 && digits[0] == '1':
		return digits[1:], nil
	default:
		return "", errors.New("Not a phone number: " + number)
	}
}

/*AreaCode extracts the area code from surrounding text*/
func AreaCode(number string) (string, error) {
	digits, err := Number(number)
	if err != nil {
		return "", err
	}
	return digits[:3], nil
}

/*Format extracts a phone number from surrounding text and formats it.*/
func Format(number string) (string, error) {
	digits, err := Number(number)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", digits[:3], digits[3:6], digits[6:10]), nil
}
