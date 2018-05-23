package accumulate

func Accumulate(items []string, f func(string) string) []string {
	result := []string{}
	for _, item := range items {
		result = append(result, f(item))
	}
	return result
}
