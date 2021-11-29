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

	for idx, name := range sortedChildren {
		if _, ok := allPlants[name]; ok {
			return &Garden{}, errors.New("duplicate name")
		}

		plants, err := handlePlantline(idx, len(children), plantLines)
		if err != nil {
			return nil, err
		}
		allPlants[name] = plants

	}
	return &Garden{plants: allPlants}, nil
}

func handlePlantline(idx int, childrenCount int, plantLines []string) ([]string, error) {
	var plants []string
	for _, plantLine := range plantLines[1:] {
		if len(plantLine) != 2*childrenCount {
			return nil, errors.New("invalid format")
		}
		for _, plant := range plantLine[2*idx : 2*idx+2] {
			val, ok := plantLookup[plant]
			if !ok {
				return nil, errors.New("plants invalid")
			}
			plants = append(plants, val)
		}
	}
	return plants, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.plants[child]

	if !ok {
		return nil, false
	}

	return plants, true
}
