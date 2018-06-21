package zebra

import "fmt"

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type House struct {
	Owner string
	Pet   string
	Color string
	Drink string
	Smoke string
}

func makeCheck(fields map[string]string) func([]House) bool {
	return func(_ []House) bool { return false }
}

func SolvePuzzle() Solution {
	var rules = [](func([]House) bool){
		// 1. There are five houses.
		func(solution []House) bool {
			return len(solution) == 5
		},
		// 2. The Englishman lives in the red house.
		makeCheck(map[string]string{"Owner": "English", "Color": "Red"}),
		// 3. The Spaniard owns the dog.
		makeCheck(map[string]string{"Owner": "Spanish", "Pet": "Dog"}),
		// 4. Coffee is drunk in the green house.
		makeCheck(map[string]string{"Drink": "Coffee", "Color": "Green"}),
		// 5. The Ukrainian drinks tea.
		// 7. The Old Gold smoker owns snails.
		makeCheck(map[string]string{"Smoke": "Old Gold", "Pet": "Snail"}),
		// 8. Kools are smoked in the yellow house.
		makeCheck(map[string]string{"Smoke": "Kools", "Color": "Yellow"}),
		// 9. Milk is drunk in the middle house.
		// 10. The Norwegian lives in the first house.
		// 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
		// 12. Kools are smoked in the house next to the house where the horse is kept.
		// 13. The Lucky Strike smoker drinks orange juice.
		makeCheck(map[string]string{"Smoke": "Luck Strike", "Drink": "Orange Juice"}),
		// 14. The Japanese smokes Parliaments.
		makeCheck(map[string]string{"Smoke": "Parliaments", "Owner": "Japanese"}),
		// 15. The Norwegian lives next to the blue house.
	}

	var values = map[string][]string{
		"Owners": []string{
			"English", "Spanish", "Ukrainian", "Norwegian", "Japanses"},
		"Pets": []string{
			"Dog", "Snail", "Fox", "Horse", "Zebra"},
		"Drink": []string{
			"Water", "Milk", "Tea", "Orange Juice", "Coffee"},
		"Smokes": []string{
			"Kools", "Old Gold", "Chesterfields", "Lucky Strike", "Parliament"},
		"Color": []string{
			"Red", "Green", "Ivory", "Yellow", "Blue"},
	}

	fmt.Println(values)
	fmt.Println(rules)

	return Solution{DrinksWater: "foo", OwnsZebra: "bar"}
}
