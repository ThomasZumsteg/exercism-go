package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

//Garden lists plants owned by students in a kindergarten garden.
type Garden map[string][]string

//plantNames are the code used to idetify each plant.
var plantNames = map[rune]string{
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

/*NewGarden creates a new kindergarten garden.*/
func NewGarden(diagram string, children []string) (*Garden, error) {
	sortedChilderen := make([]string, len(children))
	copy(sortedChilderen, children)
	sort.Strings(sortedChilderen)
	for c, child := range sortedChilderen[1:] {
		if sortedChilderen[c] == child {
			return nil, fmt.Errorf("Child is listed twice %s", child)
		}
	}

	garden := Garden{}

	if len(diagram) <= 0 || diagram[0] != '\n' {
		return nil, fmt.Errorf("Not a valid garden format")
	}

	rowLen := -1
	for _, row := range strings.Split(diagram[1:], "\n") {
		if rowLen == -1 {
			rowLen = len(row)
		} else if rowLen != len(row) {
			return nil, fmt.Errorf("Not a valid garden format: %q", diagram)
		}
		for p, plantCode := range row {
			plantName, ok := plantNames[plantCode]
			if !ok {
				return nil, fmt.Errorf("Not a valid plant code")
			}
			if len(sortedChilderen) <= p/2 {
				return nil, fmt.Errorf("Not enough childeren")
			}
			child := sortedChilderen[p/2]
			plants, ok := garden[child]
			if !ok {
				plants = []string{}
			}
			garden[child] = append(plants, plantName)
		}
	}
	return &garden, nil
}

/*Plants lists the plants owned by a child in the garden.*/
func (g Garden) Plants(child string) ([]string, bool) {
	plants, ok := g[child]
	return plants, ok
}
