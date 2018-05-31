package transpose

func Transpose(input []string) []string {
	result := []string{}
	for r, row := range input {
		for c, char := range row {
			for len(result) <= c {
				result = append(result, "")
			}
			for len(result[c]) < r {
				result[c] += " "
			}
			result[c] += string(char)
		}
	}
	return result
}
