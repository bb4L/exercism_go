package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var plantLookup = map[rune]string{
	'R': "radishes",
	'G': "grass",
	'C': "clover",
	'V': "violets",
}

type Garden struct {
	plants map[string][]string
}

func NewGarden(plantsOriginal string, children []string) (*Garden, error) {
	plantLines := strings.Split(plantsOriginal, "\n")
	if len(plantLines) != 3 {
		return &Garden{}, errors.New("garden format invalid")
	}

	allPlants := make(map[string][]string)
	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	for i, name := range sortedChildren {
		if _, ok := allPlants[name]; ok {
			return &Garden{}, errors.New("duplicate name")
		}

		for _, plantLine := range plantLines[1:] {
			if len(plantLine) != 2*len(children) {
				return &Garden{}, errors.New("invalid format")
			}
			for _, plant := range plantLine[2*i : 2*i+2] {
				if val, ok := plantLookup[plant]; ok {
					allPlants[name] = append(allPlants[name], val)
				} else {
					return &Garden{}, errors.New("plants invalid")
				}
			}
		}
	}
	return &Garden{plants: allPlants}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.plants[child]

	if !ok {
		return []string{}, false
	}

	return plants, true
}
