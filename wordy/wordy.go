package wordy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const num string = `-?\d+`

type operator func(a, b int) int

/*Answer calulates the answer to a math question.*/
func Answer(question string) (int, bool) {
	ops := strings.Join(getOps(), "|")
	parseRegex := fmt.Sprintf(`^What is (%s)((?: (?:%s) %s)+)\?$`, num, ops, num)
	match := regexp.MustCompile(parseRegex).FindStringSubmatch(question)
	if len(match) != 3 {
		return 0, false
	}

	result, _ := strconv.Atoi(match[1])
	tokenizer := fmt.Sprintf(" (%s) (%s)", ops, num)
	tokens := regexp.MustCompile(tokenizer).FindAllStringSubmatch(match[2], -1)
	for _, t := range tokens {
		op, _ := operators[t[1]]
		n, _ := strconv.Atoi(t[2])
		result = op(result, n)
	}
	return result, true
}

/*getOps gets the list of operations that are allowed.*/
func getOps() []string {
	var opsList = make([]string, 0, len(operators))
	for k := range operators {
		opsList = append(opsList, k)
	}
	return opsList
}

//operators are the valid operations
var operators = map[string]operator{
	"plus":          plus,
	"minus":         minus,
	"multiplied by": mult,
	"divided by":    div,
}

//plus adds two numbers
func plus(a, b int) int {
	return a + b
}

//minus subtracts two numbers
func minus(a, b int) int {
	return a - b
}

//mult multiples two numbers
func mult(a, b int) int {
	return a * b
}

//div divides two numbers
func div(a, b int) int {
	return a / b
}
