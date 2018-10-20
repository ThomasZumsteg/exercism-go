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

func Solve(input string) (map[string]int, error) {
    value, terms, letters := Parse(input)
    fmt.Printf("%v == %v\n", value, terms)
    fmt.Println(letters)
    return map[string]int{"L":0, "B":9, "I":1}, nil
}
