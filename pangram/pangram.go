package pangram

import "strings"

func IsPangram(s string) bool {
    letters := map[rune]bool{}
    for _, letter := range strings.ToLower(s) {
        letters[letter] = true
    }
    for _, l := range  "abcdefghijklmnopqrstuvwxyz" {
        if _, exists := letters[l]; !exists {
            return false
        }
    }
    return true
}
