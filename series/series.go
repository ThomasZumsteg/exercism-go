package slice

/*All creates all substrings of a given length*/
func All(n int, s string) []string {
	var slices = make([]string, 0)
	for n <= len(s) {
		slice, ok := First(n, s)
		if ok {
			slices = append(slices, slice)
		} else {
			return []string{}
		}
		s = s[1:]
	}
	return slices
}

/*Frist [sic] generates the first substring of a certain length*/
func Frist(n int, s string) string {
	return s[:n]
}

/*First generates the first substring of a certain length
handels cases when the substring is not possible*/
func First(n int, s string) (string, bool) {
	if len(s) < n || n <= 0 {
		return "", false
	}
	return s[:n], true
}
