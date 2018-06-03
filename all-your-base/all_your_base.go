package allyourbase

import "errors"

func ConvertToBase(from int, digits []int, to int) ([]int, error) {
	value, err := valueFromBase(from, digits)
	if err != nil {
		return nil, err
	}
	result, err := digitsFromValue(to, value)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func valueFromBase(base int, digits []int) (int, error) {
	if base < 2 {
		return 0, errors.New("input base must be >= 2")
	}
	total := 0
	for _, digit := range digits {
		if digit < 0 || base <= digit {
			return 0, errors.New("all digits must satisfy 0 <= d < input base")
		}
		total *= base
		total += digit
	}
	return total, nil
}

func digitsFromValue(base int, value int) ([]int, error) {
	if base < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	if value == 0 {
		return []int{0}, nil
	}
	digits := []int{}
	for value > 0 {
		digits = append([]int{value % base}, digits...)
		value /= base
	}
	return digits, nil
}
