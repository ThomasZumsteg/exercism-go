package say

import (
	"fmt"
	"strings"
)

//ones converts a digit to the english phrase
var ones = []string{
	"", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

//teens converts a ones digit in a teen number to the english phrase
var teens = []string{
	"ten", "eleven", "twelve", "thirteen", "fourteen",
	"fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

//tens converts a tens digits and the english phrase
var tens = []string{
	"", "ten", "twenty", "thirty", "forty",
	"fifty", "sixty", "seventy", "eighty", "ninty",
}

//powers coverts the position of a groups of digits to the english phrase
var powers = []string{
	"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion",
}

/*Say converts a number to the english phrase for the number*/
func Say(num uint64) string {
	if num == 0 {
		return "zero"
	}

	var powerGroups []uint64
	for 0 < num {
		powerGroups = append(powerGroups, num%1000)
		num /= 1000
	}

	strNum := ""
	for i := len(powerGroups) - 1; 0 <= i; i-- {
		if powerGroups[i] != 0 {
			strNum += sayPower(powerGroups[i]) + powers[i] + " "
		}
	}
	return strings.TrimSpace(strNum)
}

/*sayPower coverts a number to the english phrase, but only for numbers 1-999*/
func sayPower(num uint64) string {
	words := ""
	hundred, ten, one := (num%1000)/100, (num%100)/10, num%10
	if 0 < hundred {
		words += ones[hundred] + " hundred "
	}
	switch {
	case ten == 1:
		words += teens[one] + " "
	case one == 0:
		words += tens[ten] + " "
	case ten == 0:
		words += ones[one] + " "
	default:
		words += fmt.Sprintf("%s-%s ", tens[ten], ones[one])
	}
	return words
}
