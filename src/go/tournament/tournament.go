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
	name    string
	w, d, l int
}

func (t *Team) mp() int {
	return t.w + t.d + t.l
}

func (t *Team) score() int {
	return 3*t.w + t.d
}

func (t *Team) getString() string {
	fmtString := "%-31s| %2d | %2d | %2d | %2d | %2d\n"
	return fmt.Sprintf(fmtString, t.name, t.mp(), t.w, t.d, t.l, t.score())
}

const header = "Team                           | MP |  W |  D |  L |  P\n"

var teams = make(map[string]*Team)

// Tally returns the table sorted by points
func Tally(reader io.Reader, writer io.Writer) error {
	teams = make(map[string]*Team)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		err := addLine(line)
		if err != nil {
			return err
		}
	}

	sortedTeams := []*Team{}
	for _, team := range teams {
		sortedTeams = append(sortedTeams, team)
	}

	sort.Slice(sortedTeams, func(i, j int) bool {
		if sortedTeams[j].score() == sortedTeams[i].score() {
			return sortedTeams[i].name < sortedTeams[j].name
		}
		return sortedTeams[j].score() < sortedTeams[i].score()
	})

	io.WriteString(writer, header)

	for _, team := range sortedTeams {
		io.WriteString(writer, team.getString())
	}

	return nil
}

func addLine(line string) error {
	values := strings.Split(line, ";")

	if len(values) != 3 {
		return errors.New("Invalid line")
	}

	t1, ok := teams[values[0]]
	if !ok {
		t1 = new(Team)
		t1.name = values[0]
		teams[t1.name] = t1
	}

	t2, ok := teams[values[1]]
	if !ok {
		t2 = new(Team)
		t2.name = values[1]
		teams[t2.name] = t2
	}

	switch values[2] {
	case "win":
		t1.w++
		t2.l++

	case "draw":
		t1.d++
		t2.d++

	case "loss":
		t2.w++
		t1.l++

	default:
		return errors.New("Invalid value in result")
	}

	return nil
}
