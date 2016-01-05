package strain

//Ints is a collection of integers.
type Ints []int

//Lists is a collection of integer slices.
type Lists [][]int

//Strings is a collection of strings.
type Strings []string

/*Keep creates a new collection of ints that passes some test.*/
func (ints Ints) Keep(test func(int) bool) Ints {
	var kept Ints
	for _, i := range ints {
		if test(i) {
			kept = append(kept, i)
		}
	}
	return kept
}

/*Discard creates a new collection of ints that don't pass some test.*/
func (ints Ints) Discard(test func(int) bool) Ints {
	return ints.Keep(func(i int) bool { return !test(i) })
}

/*Keep creates a new collection of int slices that passes some test.*/
func (lists Lists) Keep(test func([]int) bool) Lists {
	var kept Lists
	for _, list := range lists {
		if test(list) {
			kept = append(kept, list)
		}
	}
	return kept
}

/*Keep creates a new collection of strings that passes some test.*/
func (strings Strings) Keep(test func(string) bool) Strings {
	var kept Strings
	for _, str := range strings {
		if test(str) {
			kept = append(kept, str)
		}
	}
	return kept
}
