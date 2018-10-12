package zebra

import (
    "strings"
    "strconv"
    // "fmt"
)

type Rule func(House) bool

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type House map[string]string

func makeHouses(attribute_map map[string][]string) []House {
    houses := []House{}
    for h := 0; h < 5; h++ {
        house := make(map[string]string)
        house["Position"] = strconv.Itoa(h)
        houses = append(houses, house)
    }
    for attr, values := range attribute_map {
        new_houses := []House{}
        for _, value := range values {
            for _, house := range houses {
                new_house := house
                house[attr] = value
                new_houses = append(new_houses, new_house)
            }
        }
        houses = new_houses
    }
    return houses
}

func makeRule(key1, value1, key2, value2, position string) Rule {
    return func(house House) bool {
        return false
    }
}

func SolvePuzzle() Solution {
	var values = map[string][]string{
		"Owners": strings.Split("English,Spanish,Ukrainian,Norwegian,Japanses", ","),
		"Pets": strings.Split("Dog,Snail,Fox,Horse,Zebra", ","),
        "Drink": strings.Split("Water,Milk,Tea,Orange Juice,Coffee", ","),
		"Smokes": strings.Split("Kools,Old Gold,Chesterfields,Lucky Strike,Parliament", ","),
		"Color": strings.Split("Red,Green,Ivory,Yellow,Blue", ","),
	}

    houses := makeHouses(values)
	var rules = [](Rule){
		// 2. The Englishman lives in the red house.
		makeRule("Owner", "English", "Color", "Red", "same"),
		// 3. The Spaniard owns the dog.
		makeRule("Owner", "Spanish", "Pet", "Dog", "same"),
		// 4. Coffee is drunk in the green house.
		makeRule("Drink", "Coffee", "Color", "Green", "same"),
		// 5. The Ukrainian drinks tea.
		makeRule("Owner", "Ukrainian", "Drink", "Tea", "same"),
        // 6. The green house is immediately to the right of the ivory house.
		makeRule("Color", "Green", "Color", "Ivory", "right"),
		// 7. The Old Gold smoker owns snails.
		makeRule("Smoke", "Old Gold", "Pet", "Snail", "same"),
		// 8. Kools are smoked in the yellow house.
		makeRule("Smoke", "Kools", "Color", "Yellow", "same"),
		// 9. Milk is drunk in the middle house.
		makeRule("Position", "middle", "Drink", "Milk", "same"),
		// 10. The Norwegian lives in the first house.
		makeRule("Position", "first", "Owner", "Norwegian", "same"),
        // 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
		makeRule("Smoke", "Chesterfields", "Pet", "Fox", "next to"),
        // 12. Kools are smoked in the house next to the house where the horse is kept.
		makeRule("Smoke", "Kools", "Pet", "Horse", "next to"),
		// 13. The Lucky Strike smoker drinks orange juice.
		makeRule("Smoke", "Luck Strike", "Drink", "Orange Juice", "same"),
		// 14. The Japanese smokes Parliaments.
		makeRule("Smoke", "Parliaments", "Owner", "Japanese", "same"),
	}
    
    for _, rule := range rules {
        
    }

	return Solution{DrinksWater: "Norwegian", OwnsZebra: "Japanese"}
}
