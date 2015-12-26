package romannumerals

import (
	"errors"
	"strings"
)

//TestVersion is the unit tests that will pass.
const TestVersion = 1

//romanToDec stores a maping between roman numerals and decimal equivalents.
var romanToDec = []struct {
	decimal int
	roman   string
}{
	{1, "I"}, {4, "IV"}, {5, "V"}, {9, "IX"},
	{10, "X"}, {40, "XL"}, {50, "L"}, {90, "XC"},
	{100, "C"}, {400, "CD"}, {500, "D"}, {900, "CM"},
	{1000, "M"},
}

/*ToRomanNumeral converts a decimal number to roman numerals
only works for numbers between 0 and 4000.*/
func ToRomanNumeral(dec int) (string, error) {
	if dec <= 0 || 4000 <= dec {
		return "", errors.New("Decimal number must be in range 0-4000")
	}
	roman := ""
	n := 0
	for i := len(romanToDec) - 1; 0 <= i; i-- {
		n, dec = divMod(dec, romanToDec[i].decimal)
		roman += strings.Repeat(romanToDec[i].roman, n)
	}
	return roman, nil
}

/*divMod calcualtes quotient and remainer*/
func divMod(numerator, divisor int) (int, int) {
	return numerator / divisor, numerator % divisor
}
