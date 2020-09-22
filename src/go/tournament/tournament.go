package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team represents a team
type Team struct {
	name                       string
	wins, draws, losses, score int
}

const header = "Team                           | MP |  W |  D |  L |  P\n"
const fmtString = "%-31s| %2d | %2d | %2d | %2d | %2d\n"

// Tally returns the table sorted by points
func Tally(reader io.Reader, writer io.Writer) error {
	var teams = make(map[string]Team)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// err := addLine(teams, line)
		values := strings.Split(line, ";")

		if len(values) != 3 {
			return errors.New("Invalid line")
		}

		t1, t2 := teams[values[0]], teams[values[1]]
		t1.name, t2.name = values[0], values[1]

		switch values[2] {
		case "win":
			t1.wins++
			t1.score += 3
			t2.losses++

		case "draw":
			t1.draws++
			t2.draws++
			t1.score++
			t2.score++

		case "loss":
			t2.wins++
			t2.score += 3
			t1.losses++

		default:
			return errors.New("Invalid value in result")
		}

		teams[t1.name], teams[t2.name] = t1, t2
	}

	sortedTeams := make([]Team, len(teams))
	count := 0
	for _, team := range teams {
		sortedTeams[count] = team
		count++
	}

	sort.Slice(sortedTeams, func(i, j int) bool {
		if sortedTeams[i].score == sortedTeams[j].score {
			return sortedTeams[i].name < sortedTeams[j].name
		}
		return sortedTeams[j].score < sortedTeams[i].score
	})

	io.WriteString(writer, header)

	for _, team := range sortedTeams {
		fmt.Fprintf(writer, fmtString, team.name, team.wins+team.draws+team.losses, team.wins, team.draws, team.losses, team.score)
	}

	return nil
}
