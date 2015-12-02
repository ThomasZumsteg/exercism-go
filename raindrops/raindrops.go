package raindrops

import "fmt"

// TestVersion the version of the tests to run
const TestVersion = 1

// rain drop sounds
var rain = make([]string, 10)

/*Convert converts a number to raindrop sounds
Adds Pling if the number is divisible by 3
     Plang if the number is divisible by 5
     Plong if the number is divisible by 7
     Otherwise returns the number as a string */
func Convert(num int) string {
	sounds := ""
	rain[3] = "Pling"
	rain[5] = "Plang"
	rain[7] = "Plong"
	for i, sound := range rain {
		if 0 < i && num%i == 0 {
			sounds += sound
		}
	}
	if sounds == "" {
		sounds = fmt.Sprintf("%d", num)
	}
	return sounds
}
