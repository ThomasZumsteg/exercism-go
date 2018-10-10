package alphametics

import (
	"fmt"
	"strings"
)

func Parse(input string) ([]string, [][]string, []string) {
	result := strings.Split(input, " == ")

	value := strings.Split(result[1], "")
	terms := [][]string{}
	letters := map[string]bool{}

	for _, l := range value {
		letters[l] = true
	}
	for _, term := range strings.Split(result[0], " + ") {
		term := strings.Split(term, "")
		terms = append(terms, term)
		for _, l := range term {
			letters[l] = true
		}
	}
	letter_list := []string{}
	for key, val := range letters {
		if val {
			letter_list = append(letter_list, key)
		}
	}
	return value, terms, letter_list
}

func Translate(word []string, letter_map map[string]int) (int, error) {
	result := 0
	for _, letter := range word {
		if digit, ok := letter_map[letter]; ok {
			result = result*10 + digit
		} else {
			return result, fmt.Errorf("Unmapped letter %s", letter)
		}
	}
	return result, nil
}

func Combo(keys []string, values []int) func() (bool, map[string]int) {
	key := keys[0]
	v := 0
	if len(keys) == 1 {
		return func() (bool, map[string]int) {
			v += 1
			if v <= len(values) {
				return true, map[string]int{key: values[v-1]}
			}
			return false, nil
		}
	} else {
		new_values := []int{}
		for i, val := range values {
			if i != v {
				new_values = append(new_values, val)
			}
		}
		gen := Combo(keys[1:], new_values)
		return func() (bool, map[string]int) {
			for true {
				if ok, combo := gen(); ok {
					combo[key] = values[v]
					return true, combo
				}
				v += 1
				if len(values) <= v {
					break
				}
				new_values = []int{}
				for i, val := range values {
					if i != v {
						new_values = append(new_values, val)
					}
				}
				gen = Combo(keys[1:], new_values)
			}
			return false, nil
		}
	}
}

func Solve(input string) (map[string]int, error) {
	value, terms, letters := Parse(input)
	gen := Combo(letters, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
Iteration:
	for ok, letter_map := gen(); ok; ok, letter_map = gen() {
		total := 0
		if digit, ok := letter_map[value[0]]; ok && digit == 0 {
			continue Iteration
		}
		for _, term := range terms {
			if value, ok := letter_map[term[0]]; ok && value == 0 {
				continue Iteration
			}
			result, _ := Translate(term, letter_map)
			total += result
		}
		compare, _ := Translate(value, letter_map)
		if compare == total {
			return letter_map, nil
		}
	}
	return nil, fmt.Errorf("No valid solution")
}
