package forth

import (
	"errors"
	"strconv"
	"strings"
)

func Forth(lines []string) ([]int, error) {
	stack := []int{}
	words := make(map[string][]string)
	for _, line := range lines {
		fields := strings.Fields(line)
		if 3 <= len(fields) && fields[0] == ":" && fields[len(fields)-1] == ";" {
			word, definition := fields[1], fields[2:len(fields)-1]
			if _, err := strconv.Atoi(word); err == nil {
				return []int(nil), errors.New("Definition cannot be a number")
			}
			words[strings.ToLower(word)] = definition
			continue
		}
		for i := 0; i < len(fields); i++ {
			tok := strings.ToLower(fields[i])
			if num, err := strconv.Atoi(tok); err == nil {
				stack = append(stack, num)
			} else if val, ok := words[tok]; ok {
				head, tail := fields[:i+1], fields[i+1:]
				fields = append(head, val...)
				fields = append(fields, tail...)
			} else {
				if stack, err = builtIns(tok, stack); err != nil {
					return []int(nil), err
				}
			}
		}
	}
	return stack, nil
}

func builtIns(tok string, stack []int) ([]int, error) {
	switch tok {
	case "+", "-", "*", "/":
		if len(stack) < 2 {
			return []int(nil), errors.New("Parse error")
		}
		l := len(stack)
		b, a := stack[l-1], stack[l-2]
		stack = stack[:l-2]
		switch tok {
		case "+":
			stack = append(stack, a+b)
		case "-":
			stack = append(stack, a-b)
		case "*":
			stack = append(stack, a*b)
		case "/":
			if b == 0 {
				return []int(nil), errors.New("Divide by zero")
			}
			stack = append(stack, a/b)
		}
	case "dup":
		if len(stack) < 1 {
			return []int(nil), errors.New("Parse error")
		}
		l := len(stack)
		stack = append(stack, stack[l-1])
	case "drop":
		if len(stack) < 1 {
			return []int(nil), errors.New("Parse error")
		}
		stack = stack[:len(stack)-1]
	case "swap":
		if len(stack) < 2 {
			return []int(nil), errors.New("Parse error")
		}
		l := len(stack)
		stack[l-1], stack[l-2] = stack[l-2], stack[l-1]
	case "over":
		if len(stack) < 2 {
			return []int(nil), errors.New("Parse error")
		}
		l := len(stack)
		stack = append(stack, stack[l-2])
	default:
		return []int(nil), errors.New("Unrcognized Token")
	}
	return stack, nil
}
