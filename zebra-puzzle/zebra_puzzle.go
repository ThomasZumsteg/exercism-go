package zebra

import (
    "strings"
    "sort"
    "strconv"
    // "fmt"
)

type House map[string]string
type Rule func([]House) bool

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type ByHash []string

func (a ByHash) Len() int { return len(a) }
func (a ByHash) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHash) Less(i, j int) bool { return a[i] < a[j] }


func Permutations(items []string) func() []string {
    n_items := len(items)
    sort.Sort(ByHash(items))
    return func() []string {
        for k := n_items - 1; k <= 0; k-- {
            if items[k] < items[k+1] {
                items[k], items[k+1] = items[k+1], items[k]
                ret_items := make([]string, n_items)
                copy(ret_items, items)
                return ret_items
            }
        }
        return nil;
    }
}


func MakeSolutions(attribute_map map[string][]string) [][]House {
    solutions := [][]House{[]House{
        House{"Position": "1"},
        House{"Position": "2"},
        House{"Position": "3"},
        House{"Position": "4"},
        House{"Position": "5"}}}
    for attr, values := range attribute_map {
        new_solutions := [][]House{}
        for _, houses := range solutions {
            gen := Permutations(values)
            for perm := gen(); perm != nil; perm = gen() {
                new_iter := make([]House, 5)
                copy(new_iter, houses)
                for i := 0; i < len(values); i++ {
                    new_iter[i][attr] = values[i]
                }
                new_solutions = append(new_solutions, new_iter)
            }
        }
        solutions = new_solutions
    }
    return solutions
}

func makeRule(key1, value1, key2, value2, position string) Rule {
    return func(houses []House) bool {
        var pos1, pos2 int
        for _, house := range houses {
            if house[key1] == value1 {
                pos1, _ = strconv.Atoi(house["position"])
            }
            if house[key2] == value2 {
                pos2, _ = strconv.Atoi(house["position"])
            }
        }
        switch position {
        case "same": return pos1 == pos2
        case "right": return pos1 + 1 == pos2
        case "left": return pos1 == 1 + pos2
        case "next to": return (pos1 + 1 == pos2) || (pos1 == pos2 + 1)
        }
        panic("Not a valid rule")
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

    solutions :=  MakeSolutions(values)
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
        new_solutions := [][]House{}
        for _, sol := range solutions {
            if rule(sol) {
                new_solutions = append(new_solutions, sol)
            }
        }
        solutions = new_solutions
    }

	return Solution{DrinksWater: "Norwegian", OwnsZebra: "Japanese"}
}
