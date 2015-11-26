package raindrops

import "fmt"

// TestVersion the version of the tests to run
const TestVersion = 1

/* Convert converts a number to raindrop sounds
Adds Pling if the number is divisible by 3
     Plang if the number is divisible by 5
     Plong if the number is divisible by 7
     Otherwise returns the number as a string */
func Convert(num int) string {
    sounds := ""
    if num % 3 == 0 {
        sounds += "Pling"
    }
    if num % 5 == 0 {
        sounds += "Plang"
    }
    if num % 7 == 0 {
        sounds += "Plong"
    }
    if sounds == "" {
        sounds = fmt.Sprintf("%d", num)
    }
    return sounds
}

