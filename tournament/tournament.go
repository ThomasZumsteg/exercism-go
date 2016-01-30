package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

//TestVersion is the verion of the unit tests that this will pass
const TestVersion = 2

//scoreBoard keeps track of the score
type scoreBoard map[string]*team

//team is the win/loss record of the team
type team struct {
	name                            string
	played, win, loss, draw, points int
}

/*Tally counts up the wins and losses for a sports season.*/
func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	board := make(scoreBoard)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if err := board.addGame(line); err != nil {
			return err
		}
	}

	teams := board.getTeams()
	sort.Sort(byScore(teams))
	header := "Team                           | MP |  W |  D |  L |  P\n"
	io.WriteString(writer, header)
	for _, team := range teams {
		io.WriteString(writer, team.String()+"\n")
	}
	return nil
}

/*addGame validates and adds a game to the scoreboard.*/
func (b scoreBoard) addGame(game string) error {
	fields := strings.Split(game, ";")
	if len(fields) != 3 {
		return fmt.Errorf("Game not properly formatted: %s", game)
	}
	home, homeOk := b[fields[0]]
	away, awayOk := b[fields[1]]
	if !homeOk {
		home = &team{name: fields[0]}
		b[fields[0]] = home
	}
	if !awayOk {
		away = &team{name: fields[1]}
		b[fields[1]] = away
	}
	switch fields[2] {
	case "win":
		home.addWin()
		away.addLoss()
	case "loss":
		home.addLoss()
		away.addWin()
	case "draw":
		home.addDraw()
		away.addDraw()
	default:
		return fmt.Errorf("Unknown win condition: %s", game)
	}
	return nil
}

/*getTeams converts the scoreboard into a list.*/
func (b scoreBoard) getTeams() []team {
	var teams []team
	for _, team := range b {
		teams = append(teams, *team)
	}
	return teams
}

//byScore sorts the teams by points, then by wins, then aphabetically
type byScore []team
func (t byScore) Len() int      { return len(t) }
func (t byScore) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t byScore) Less(i, j int) bool {
	if t[i].points != t[j].points {
		return t[i].points > t[j].points
	} else if t[i].win != t[j].win {
		return t[i].win > t[j].win
	}
	return t[i].name < t[j].name
}

/*addWin adds a win to the team.*/
func (t *team) addWin() {
	t.played++
	t.win++
	t.points += 3
}

/*addLoss adds a loss to the team.*/
func (t *team) addLoss() {
	t.played++
	t.loss++
}

/*addDraw adds a draw to the team.*/
func (t *team) addDraw() {
	t.played++
	t.draw++
	t.points++
}

/*String gets the record of the team.*/
func (t *team) String() string {
	fmtString := "%-31s| %2d | %2d | %2d | %2d | %2d"
	return fmt.Sprintf(fmtString, t.name, t.played, t.win, t.draw, t.loss, t.points)
}
