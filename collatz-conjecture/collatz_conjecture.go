package collatzconjecture

import "fmt"

func CollatzConjecture(num int) (i int, err error) {
	if num <= 0 {
		err = fmt.Errorf("Integers must be greater than zero: %d", num)
		return
	}
	for i = 0; num != 1; i++ {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
	}
	return
}
