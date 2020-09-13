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
	return string(charset[random.Intn(len(charset))]) + string(charset[random.Intn(len(charset))]) + strconv.Itoa(random.Intn(10)) + strconv.Itoa(random.Intn(10)) + strconv.Itoa(random.Intn(10))
}

// Name get the name of the robot
func (r *Robot) Name() (string, error) {
	if len(names) == 26*26*10*10*10*10 {
		return "", errors.New("Too many robots")
	}
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
	removeName(r.name)
	r.name = ""
}
