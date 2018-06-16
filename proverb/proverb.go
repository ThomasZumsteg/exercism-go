package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	result := []string{}
	if len(rhyme) <= 0 {
		return result
	}

	for i := 0; i+1 < len(rhyme); i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		result = append(result, line)
	}
	result = append(result, fmt.Sprintf(
		"And all for the want of a %s.",
		rhyme[0]))
	return result
}
