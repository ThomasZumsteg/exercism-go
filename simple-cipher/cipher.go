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
type cipher []int

/*shiftLetter encodes a letter based on some function.*/
func (c cipher) shiftLetters(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(s)
	}
	return shiftedText
}

/*Encode encrypts a message.*/
func (c *cipher) Encode(plainText string) string {
	return c.shiftLetters(plainText, func(a, b int) int { return a + b })
}

/*Decode decrypts a message.*/
func (c *cipher) Decode(cipherText string) string {
	return c.shiftLetters(cipherText, func(a, b int) int { return a - b })
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
	c := cipher([]int{shift})
	return &c
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
	c := cipher(shiftInt)
	return &c
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
