package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	result := 0
	for _, birdCount := range birdsPerDay {
		result += birdCount
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	result := 0
	for i := (week - 1) * 7; i < week*7; i++ {
		result += birdsPerDay[i]
	}
	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx, val := range birdsPerDay {
		if idx%2 == 0 {
			birdsPerDay[idx] = val + 1
		}
	}
	return birdsPerDay
}
