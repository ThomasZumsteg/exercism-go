package cipher

import (
	"unicode"
)

//Cipher encrypts and decrypts a string.
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

//cipher holds the key used to encipher messages.
type cipher struct {
	shiftDist []int
	index     int
}

/*shiftLetter encodes a letter based on some function.*/
func (c *cipher) shiftLetter(char rune, shift func(int, int) int) string {
	if !unicode.IsLetter(char) {
		return ""
	}
	shiftDist := c.shiftDist[c.index%len(c.shiftDist)]
	c.index++
	s := shift(int(unicode.ToLower(char)), shiftDist)
	switch {
	case s < 'a':
		s += 'z' - 'a' + 1
	case 'z' < s:
		s -= 'z' - 'a' + 1
	}
	return string(s)
}

/*Encode encrypts a message.*/
func (c *cipher) Encode(plainText string) string {
	cipherText := ""
	for _, char := range plainText {
		cipherText += c.shiftLetter(char, func(a, b int) int { return a + b })
	}
	c.index = 0
	return cipherText
}

/*Decode decrypts a message.*/
func (c *cipher) Decode(cipherText string) string {
	plainText := ""
	for _, char := range cipherText {
		plainText += c.shiftLetter(char, func(a, b int) int { return a - b })
	}
	c.index = 0
	return plainText
}

/*NewCaesar creates a new Caesar shift cipher.*/
func NewCaesar() Cipher {
	return NewShift(3)
}

/*NewShift creates a new Shift cipher.*/
func NewShift(shift int) Cipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	return &cipher{shiftDist: []int{shift}}
}

/*NewVigenere creates a new Vigenere cipher.*/
func NewVigenere(shift string) Cipher {
	switch {
	case shift == "":
		return nil
	case all(shift, func(r rune) bool { return r == 'a' }):
		return nil
	}
	shiftInt := make([]int, len(shift))
	for i, c := range shift {
		if !unicode.IsLower(c) || !unicode.IsLetter(c) {
			return nil
		}
		shiftInt[i] = int(c - 'a')
	}
	return &cipher{shiftDist: shiftInt}
}

/*all check if all items in a string pass some function.*/
func all(items string, test func(rune) bool) bool {
	for _, r := range items {
		if !test(r) {
			return false
		}
	}
	return true
}
