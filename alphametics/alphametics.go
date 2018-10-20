package alphametics

import (
    "strings"
    "fmt"
)

func Parse(input string) ([]string, [][]string, map[string]bool) {
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
    return value, terms, letters
}

func Translate(word []string, letter_map map[string]int) (int, error) {
    result := 0
    for _, letter := range word {
        if digit, ok := letter_map[letter]; ok {
            result = result * 10 + digit
        } else {
            return result, fmt.Errorf("Unmapped letter %s", letter)
        }
    }
    return result, nil
}

func Solve(input string) (map[string]int, error) {
    value, terms, letters := Parse(input)
    letter_map := map[string]int{}
    for i, letter := range letters {
        letter_map[letter] = i
    }
    for true {
        total := 0
        for _, term := range terms {
            result, _ := Translate(term, letter_map)
            total += result
        }
        compare, _ := Translate(value, letter_map)
        if compare == total {
            return letter_map, nil
        }
        letter_map[letters[0]] += 1
        for _, letter := range letter {
        }
    }
    return nil, fmt.Errorf("No valid solution")
}
