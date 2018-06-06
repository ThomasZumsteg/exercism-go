package forth

import (
    "strings"
    "strconv"
    "errors"
)

func Forth(lines []string) ([]int, error) {
    stack := []int{}
    for _, line := range lines {
        for _, tok := range strings.Fields(line) {
            if i, err := strconv.Atoi(tok); err == nil {
                stack = append(stack, i)
                continue
            }
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
                    stack = append(stack, a + b)
                case "-":
                    stack = append(stack, a - b)
                case "*":
                    stack = append(stack, a * b)
                case "/":
                    if b == 0 {
                        return []int(nil), errors.New("Divide by zero")
                    }
                    stack = append(stack, a / b)
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
            }
        }
    }
    return stack, nil
}
