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

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var names = make(map[string]int)

func nameExists(value string) bool {
	_, ok := names[value]
	return ok
}

func removeName(val string) {
	// delete(names, val)
}

func createName() string {
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}

// Name get the name of the robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(names) == 26*26*10*10*10*10 {
			return "", errors.New("Too many robots")
		}
		name := createName()
		for nameExists(name) {
			name = createName()
		}
		names[name] = 1
		r.name = name
	}
	return r.name, nil
}

// Reset resets the name
func (r *Robot) Reset() {
	removeName(r.name)
	r.name = ""
}
