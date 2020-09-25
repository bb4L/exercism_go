package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot type for a robot
type Robot struct {
	name string
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var names = make(map[string]bool)

func nameExists(value string) bool {
	_, ok := names[value]
	return ok
}

func createName() string {
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}

// Name get the name of the robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if len(names) == 26*26*10*10*10 {

		return "", errors.New("Too many robots")
	}
	r.name = createName()
	for names[r.name] {
		r.name = createName()
	}
	names[r.name] = true
	return r.name, nil
}

// Reset resets the name
func (r *Robot) Reset() {
	r.name = ""
}
