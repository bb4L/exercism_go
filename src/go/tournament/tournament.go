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
	name                string
	wins, draws, losses int
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

		t1, ok := teams[values[0]]
		if !ok {
			t1.name = values[0]
		}

		t2, ok := teams[values[1]]
		if !ok {
			t2.name = values[1]
		}

		switch values[2] {
		case "win":
			t1.wins++
			t2.losses++

		case "draw":
			t1.draws++
			t2.draws++

		case "loss":
			t2.wins++
			t1.losses++

		default:
			return errors.New("Invalid value in result")
		}

		teams[t1.name] = t1
		teams[t2.name] = t2

	}

	sortedTeams := []Team{}
	for _, team := range teams {
		sortedTeams = append(sortedTeams, team)
	}

	sort.Slice(sortedTeams, func(i, j int) bool {
		scoreJ := 3*sortedTeams[j].wins + sortedTeams[j].draws
		scoreI := 3*sortedTeams[i].wins + sortedTeams[i].draws
		if scoreI == scoreJ {
			return sortedTeams[i].name < sortedTeams[j].name
		}
		return scoreJ < scoreI
	})

	io.WriteString(writer, header)

	for _, team := range sortedTeams {
		io.WriteString(writer, fmt.Sprintf(fmtString, team.name, team.wins+team.draws+team.losses, team.wins, team.draws, team.losses, 3*team.wins+team.draws))
	}

	return nil
}
