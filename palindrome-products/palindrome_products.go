package palindrome

import (
	"errors"
	"strconv"
)

//Product holds a palindrome number and its factors
type Product struct {
	Product        int
	Factorizations [][2]int
}

/*Products finds the minimum and maximum palindrome products in a range of numbers
and checks that there is a palindrome and that the bonds are correct*/
func Products(fmin, fmax int) (Product, Product, error) {
	// if fmax < 0 || fmin < 0 {
	// 	return Product{}, Product{}, errors.New("Negative limits")
	// }
	if fmax < fmin {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}
	var palindromes []Product
	for factorSmall := fmin; factorSmall <= fmax; factorSmall++ {
		for factorLarge := factorSmall; factorLarge <= fmax; factorLarge++ {
			product := factorSmall * factorLarge
			digits := strconv.Itoa(product)
			if isPalindrome(digits) {
				p := Product{product, [][2]int{{factorSmall, factorLarge}}}
				palindromes = insert(palindromes, p)
			}
		}
	}
	if len(palindromes) <= 0 {
		return Product{}, Product{}, errors.New("No palindromes")
	}
	return palindromes[0], palindromes[len(palindromes)-1], nil
}

/*isPalindrome determines if a string is a palindrome*/
func isPalindrome(s string) bool {
	l := len(s) - 1
	for i := 0; i < l-i; i++ {
		if s[i] != s[l-i] {
			return false
		}
	}
	return true
}

/*insert adds a palindrome number into a sorted list of palindrome number*/
func insert(list []Product, p Product) []Product {
	for i := 0; i < len(list); i++ {
		if p.Product < list[i].Product {
			return append(list[:i], append([]Product{p}, list[i:]...)...)
		} else if list[i].Product == p.Product {
			list[i].Factorizations = append(list[i].Factorizations, p.Factorizations...)
			return list
		}
	}
	return append(list, p)
}
