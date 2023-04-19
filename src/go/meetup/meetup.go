package meetup

import "time"

// WeekSchedule data structure to store a weekday
type WeekSchedule int

// First indicates looking for the first
var First = WeekSchedule(1)

// Second indicates looking for the second
var Second = WeekSchedule(2)

// Third indicates looking for the third
var Third = WeekSchedule(3)

// Fourth indicates looking for the fourth
var Fourth = WeekSchedule(4)

// Last indicates lookinf for the last
var Last = WeekSchedule(-1)

// Teenth indicates looking for a value in the x-'teen'
var Teenth = WeekSchedule(10)

// Day return the day for the given input
func Day(schedule WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	startValue := 1
	countTo := 1
	changeValue := 1

	switch schedule {
	case First, Second, Third, Fourth:
		countTo = int(schedule)

	case Teenth:
		startValue = 13

	case Last:
		startValue = 31
		changeValue = -1
	}

	count := 0
	for i := startValue; ; i = i + changeValue {
		dateToTest := time.Date(year, month, i, 0, 0, 0, 0, time.UTC)
		if dateToTest.Month() != month {
			continue
		}
		if dateToTest.Weekday() == weekday {
			count++
			if countTo == count {
				return i
			}
		}
	}
}
