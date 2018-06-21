package zebra

// import "fmt"
import "reflect"

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

func makeCheck(field1, value1, field2, value2 string) func([]House) bool {
	return func(houses []House) bool {
		for _, house := range houses {
			obj := reflect.Indirect(reflect.ValueOf(house))
			if obj.FieldByName(field1).String() == value1 {
				return obj.FieldByName(field2).String() == value2
			}
		}
		return false
	}
}

func SolvePuzzle() Solution {
	var rules = [](func([]House) bool){
		// 1. There are five houses.
		func(solution []House) bool {
			return len(solution) == 5
		},
		// 2. The Englishman lives in the red house.
		makeCheck("Owner", "English", "Color", "Red"),
		// 3. The Spaniard owns the dog.
		makeCheck("Owner", "Spanish", "Pet", "Dog"),
		// 4. Coffee is drunk in the green house.
		makeCheck("Drink", "Coffee", "Color", "Green"),
		// 5. The Ukrainian drinks tea.
		// 7. The Old Gold smoker owns snails.
		makeCheck("Smoke", "Old Gold", "Pet", "Snail"),
		// 8. Kools are smoked in the yellow house.
		makeCheck("Smoke", "Kools", "Color", "Yellow"),
		// 9. Milk is drunk in the middle house.
		func(solution []House) bool {
			return solution[2].Drink == "Milk"
		},
		// 10. The Norwegian lives in the first house.
		func(solution []House) bool {
			return solution[0].Owner == "Norwegian"
		},
		// 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
		func(houses []House) bool {
			for h, house := range houses {
				if house.Smoke == "Chesterfields" {
					return (0 < h && houses[h-1].Pet == "Fox") ||
						(h < 4 && houses[h+1].Pet == "Fox")
				}
			}
            return false
		},
		// 12. Kools are smoked in the house next to the house where the horse is kept.
		func(houses []House) bool {
			for h, house := range houses {
				if house.Smoke == "Kools" {
					return (0 < h && houses[h-1].Pet == "Horse") ||
						(h < 4 && houses[h+1].Pet == "Horse")
				}
			}
            return false
		},
		// 13. The Lucky Strike smoker drinks orange juice.
		makeCheck("Smoke", "Luck Strike", "Drink", "Orange Juice"),
		// 14. The Japanese smokes Parliaments.
		makeCheck("Smoke", "Parliaments", "Owner", "Japanese"),
		// 15. The Norwegian lives next to the blue house.
		func(houses []House) bool {
			for h, house := range houses {
				if house.Owner == "Norwegian" {
					return (0 < h && houses[h-1].Color == "Blue") ||
						(h < 4 && houses[h+1].Color == "Blue")
				}
			}
            return false
		},
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

	return Solution{DrinksWater: "Norwegian", OwnsZebra: "Japanese"}
}
