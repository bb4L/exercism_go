package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// Robot type for a robot
type Robot struct {
	name string
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

var names = make(map[string]int)

func nameExists(value string) bool {
	_, ok := names[value]
	return ok
}

func createName() string {
	return string(65+random.Intn(26)) + string(65+random.Intn(26)) + strconv.Itoa(random.Intn(10)) + strconv.Itoa(random.Intn(10)) + strconv.Itoa(random.Intn(10))
}

// Name get the name of the robot
func (r *Robot) Name() (string, error) {
	if len(names) == 26*26*10*10*10*10 {
		return "", errors.New("Too many robots")
	}

	// this is a special case to handle
	if r.name == "" {
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
	r.name = ""
}
