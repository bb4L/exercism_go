package clock

import (
	"fmt"
)

// Clock representing a clock
type Clock struct {
	hour   int
	minute int
}

func (c *Clock) correctValues() {
	if c.minute > 59 || c.minute < 0 {
		c.hour += c.minute / 60
		c.minute = c.minute % 60

		if c.minute < 0 {
			c.minute += 60
			c.hour--
		}
	}

	c.hour = c.hour % 24
	if c.hour < 0 {
		c.hour += 24
	}
}

// New create a new Clock
func New(hour, minute int) Clock {
	c := Clock{hour, minute}
	c.correctValues()
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract subtract minutes
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
