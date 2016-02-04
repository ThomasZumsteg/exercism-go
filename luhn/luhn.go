package luhn

import (
	"errors"
	"strconv"
	"unicode"
)

/*Valid determines if a luhn code is correct.*/
func Valid(code string) bool {
	t, err := checkSum(code)
	return err == nil && t%10 == 0
}

/*AddCheck creates a valid luhn code.*/
func AddCheck(code string) string {
	finalDigit := "0"
	t, _ := checkSum(code + "0")
	if t%10 != 0 {
		finalDigit = strconv.Itoa(10 - (t % 10))
	}
	if all(digitsAreFullWidth, code) {
		// Added digit should be full width if all the others are
		finalDigit = string(int(finalDigit[0]) + int('０') - '0')
	}
	return code + finalDigit
}

/*all checks that a condition is true of all elements.*/
func all(condition func(rune) bool, items string) bool {
	for _, char := range items {
		if !condition(char) {
			return false
		}
	}
	return true
}

/*digitsAreFullWidth checks if a rune is not a digit or full width.*/
func digitsAreFullWidth(char rune) bool {
	digit := int(char) - int('０')
	return !unicode.IsDigit(char) || 0 <= digit && digit < 10
}

/*checkSum computes the check sum of a luhn code*/
func checkSum(code string) (int, error) {
	numDigits := 0
	total := 0
	for i := len(code) - 1; 0 <= i; i-- {
		n, err := strconv.Atoi(string(code[i]))
		switch {
		case err != nil:
			continue
		case numDigits%2 == 0:
			total += n
		case n <= 4:
			total += 2 * n
		case n <= 9:
			total += 2*n - 9
		}
		numDigits++
	}
	if numDigits <= 0 {
		return 0, errors.New("There are no digits in \"" + code + "\"")
	}
	return total, nil
}
