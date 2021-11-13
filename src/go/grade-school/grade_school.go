package school

import (
	"sort"
)

// Define the Grade and School types here.
type Grade struct {
	grade   int
	studens []string
}
type School struct {
	grades []Grade
}

func New() *School {
	return &School{grades: make([]Grade, 9)}
}

func (s *School) Add(student string, grade int) {
	s.grades[grade-1].studens = append(s.grades[grade-1].studens, student)
	s.grades[grade-1].grade = grade
	sort.Strings(s.grades[grade-1].studens)
}

func (s *School) Grade(level int) []string {
	return s.grades[level-1].studens
}

func (s *School) Enrollment() (result []Grade) {
	for _, grade := range s.grades {
		if grade.grade > 0 {
			result = append(result, grade)
		}
	}
	return
}
