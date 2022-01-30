package school

import (
	"sort"
)

// Grade holds the information for a given grade
type Grade struct {
	grade   int
	studens []string
}

// School stores all grades
type School struct {
	grades []Grade
}

// New creates a new School
func New() *School {
	return &School{grades: make([]Grade, 9)}
}

// Add adds a student to the given grade
func (s *School) Add(student string, grade int) {
	s.grades[grade-1].studens = append(s.grades[grade-1].studens, student)
	s.grades[grade-1].grade = grade
	sort.Strings(s.grades[grade-1].studens)
}

// Grade return all students with a given grade
func (s *School) Grade(level int) []string {
	return s.grades[level-1].studens
}

// Enrollment return all grades greater than 0
func (s *School) Enrollment() (result []Grade) {
	for _, grade := range s.grades {
		if grade.grade > 0 {
			result = append(result, grade)
		}
	}
	return
}
