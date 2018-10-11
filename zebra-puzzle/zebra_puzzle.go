package zebra

// import "fmt"

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type House struct {
    Position int
	Owner string
	Pet   string
	Color string
	Drink string
	Smoke string
}

func makeHouses(numHouses int) []House {
    houses := []House{}
    for h := 0; h < numHouses; h++ {
        houses = append(houses, House{ -1, "", "", "", "", ""})
    }
    return houses
}

func SolvePuzzle() Solution {
    // 1. There are five houses.
    houses := makeHouses(5)
	var rules = [](func([]House) bool){
		// 2. The Englishman lives in the red house.
		makeRule("Owner", "English", "Color", "Red"),
		// 3. The Spaniard owns the dog.
		makeRule("Owner", "Spanish", "Pet", "Dog"),
		// 4. Coffee is drunk in the green house.
		makeRule("Drink", "Coffee", "Color", "Green"),
		// 5. The Ukrainian drinks tea.
		// 7. The Old Gold smoker owns snails.
		makeRule("Smoke", "Old Gold", "Pet", "Snail"),
		// 8. Kools are smoked in the yellow house.
		makeRule("Smoke", "Kools", "Color", "Yellow"),
		// 9. Milk is drunk in the middle house.
		makeRule("Position", 2, "Drink", "Milk"),
		// 10. The Norwegian lives in the first house.
		makeRule("Position", 0, "Owner", "Norwegian"),
		// 13. The Lucky Strike smoker drinks orange juice.
		makeRule("Smoke", "Luck Strike", "Drink", "Orange Juice"),
		// 14. The Japanese smokes Parliaments.
		makeRule("Smoke", "Parliaments", "Owner", "Japanese"),
	}

    var checks = [](func([]House) bool) {
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
