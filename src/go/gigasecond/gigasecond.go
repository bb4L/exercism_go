package gigasecond

import (
	"math"
	"time"
)

var gigasecond = time.Duration(math.Pow(math.Pow(10, 9), 2))

// AddGigasecond add a gigasecond to the given time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
