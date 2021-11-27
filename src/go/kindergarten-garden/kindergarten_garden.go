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

	var err error

	allPlants := make(map[string][]string)
	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	for idx, name := range sortedChildren {
		if _, ok := allPlants[name]; ok {
			return &Garden{}, errors.New("duplicate name")
		}

		allPlants, err = handlePlantline(idx, allPlants, children, name, plantLines)
		if err != nil {
			return nil, err
		}
	}
	return &Garden{plants: allPlants}, nil
}

func handlePlantline(i int, allPlants map[string][]string, children []string, name string, plantLines []string) (map[string][]string, error) {
	for _, plantLine := range plantLines[1:] {
		if len(plantLine) != 2*len(children) {
			return nil, errors.New("invalid format")
		}
		for _, plant := range plantLine[2*i : 2*i+2] {
			val, ok := plantLookup[plant]
			if !ok {
				return nil, errors.New("plants invalid")
			}
			allPlants[name] = append(allPlants[name], val)
		}
	}
	return allPlants, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.plants[child]

	if !ok {
		return nil, false
	}

	return plants, true
}
