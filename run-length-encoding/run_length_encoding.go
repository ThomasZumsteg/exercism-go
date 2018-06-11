package encode

import "strconv"

func RunLengthEncode(input string) (result string) {
    if input == "" {
        return
    }
    last := rune(input[0])
    count := 0
    for _, char := range input {
        if char == last {
            count++
        } else {
            if 1 < count {
                result += strconv.Itoa(count)
            }
            result += string(last)
            count = 1
        }
        last = char
    }
    if 1 < count {
        result += strconv.Itoa(count)
    }
    result += string(last)
    return
}

func RunLengthDecode(input string) (result string) {
    count := 0
    for _, char := range input {
        if d, err := strconv.Atoi(string(char)); err == nil {
            count *= 10
            count += d
        } else {
            if count == 0 {
                count = 1
            }
            for c := 0; c < count; c++ {
                result += string(char)
            }
            count = 0
        }
    }
    return
}
