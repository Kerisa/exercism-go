package clock

import (
	"fmt"
	"math"
)

// Clock type represent a clock
type Clock struct {
	hour int
	minute int
}

// New creates a new clock
func New(hour, minute int) Clock {
	hour += int(math.Floor(float64(minute) / 60))
	hour = (hour % 24 + 24) % 24
	minute = (minute % 60 + 60) % 60
	return Clock{hour, minute}
}

// String converts a clock into string
func (c *Clock)String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds a duration of minutes and return the new clock
func (c Clock)Add(minutes int) Clock {
	return New(c.hour, c.minute + minutes)
}

// Subtract minus a duration of minutes and return the new clock
func (c Clock)Subtract(minutes int) Clock {
	return c.Add(-minutes)
}