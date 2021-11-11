package school

import (
	"sort"
)

// Define the Grade and School types here.
type Grade struct {
	grade    int
	students []string
}
type School struct {
	grades []Grade
}

func New() *School {
	return &School{}
}

func (s *School) Add(student string, grade int) {
	for idx, val := range s.grades {
		if val.grade == grade {
			val.students = append(val.students, student)
			sort.Strings(val.students)
			s.grades[idx] = val
			return
		}
	}
	s.grades = append(s.grades, Grade{grade, []string{student}})

}

func (s *School) Grade(level int) []string {
	for _, val := range s.grades {
		if val.grade == level {
			return val.students
		}
	}
	return []string{}
}

func (s *School) Enrollment() []Grade {
	sort.Slice(s.grades, func(i, j int) bool {
		return s.grades[i].grade < s.grades[j].grade
	})
	return s.grades
}
