package allergies

/*allergieList is an ordered list of things to be allergic to
it corrosponds to the bits in an integer code*/
var allergieList = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

/*Allergies determines all things things that person is allergic to
based on a code*/
func Allergies(code int) []string {
	var allergies []string
	for i, allergy := range allergieList {
		if 0 != (code & (1 << uint(i))) {
			allergies = append(allergies, allergy)
		}
	}
	return allergies
}

/*AllergicTo determines if a person is allergic to an item*/
func AllergicTo(code int, allergen string) bool {
	for i, allergy := range allergieList {
		if allergy == allergen && 0 != (code&(1<<uint(i))) {
			return true
		}
	}
	return false
}
