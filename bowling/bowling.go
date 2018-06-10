package bowling

import "fmt"

type Game struct {
	throws []int
}

func NewGame() *Game {
	return &Game{[]int{}}
}

func (g *Game) Roll(pins int) error {
	if !g.more() {
		return fmt.Errorf("Game is done: %v", g.throws)
	} else if pins < 0 || 10 < pins ||
		(len(g.throws)%2 == 1 && 10 < g.throws[len(g.throws)-1]+pins) {
		return fmt.Errorf("Invalid number of pins: %v, %d", g.throws, pins)
	}
	g.throws = append(g.throws, pins)
	if pins == 10 {
		g.throws = append(g.throws, 0)
	}
	return nil
}

func (g *Game) Score() (int, error) {
	total := 0
	if g.more() {
		return 0, fmt.Errorf("Game is not done yet: %v", g.throws)
	}
	for t := 0; t < 20; t += 2 {
		first, second := g.throws[t], g.throws[t+1]
		total += first + second
		if first+second == 10 {
			total += g.throws[t+2]
		}
		if first == 10 {
			if g.throws[t+2] == 10 {
				total += g.throws[t+4]
			} else {
				total += g.throws[t+3]
			}
		}
	}
	return total, nil
}

func (g *Game) more() bool {
	throws := len(g.throws)
	switch {
	case 0 <= throws && throws < 20:
		return true
	case throws == 20:
		return g.throws[18]+g.throws[19] == 10
	case throws == 21:
		return g.throws[18] == 10
	case throws == 22 && g.throws[18] == 10 && g.throws[20] == 10:
		return true
	default:
		return false
	}
}
