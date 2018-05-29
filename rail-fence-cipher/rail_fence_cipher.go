package railfence

import (
	"strings"
)

func Encode(plainText string, numRails int) string {
	length := len(plainText)
	cipherText := make([]string, length)
	for to, from := range MakeCipher(len(plainText), numRails) {
		cipherText[to] = string(plainText[from])
	}
	return strings.Join(cipherText, "")
}

func Decode(cipherText string, numRails int) string {
	length := len(cipherText)
	plainText := make([]string, length)
	for from, to := range MakeCipher(len(cipherText), numRails) {
		plainText[to] = string(cipherText[from])
	}
	return strings.Join(plainText, "")
}

func MakeCipher(items int, numRails int) []int {
	var rails [][]int
	rail, dRail := 0, 1
	for i := 0; i < items; i++ {
		if rail <= numRails {
			rails = append(rails, []int{})
		}
		rails[rail] = append(rails[rail], i)
		if rail+dRail < 0 || numRails <= rail+dRail {
			dRail = -dRail
		}
		rail += dRail
	}
	result := []int{}
	for _, slice := range rails {
		result = append(result, slice...)
	}
	return result
}
