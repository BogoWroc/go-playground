package clock

import (
	"fmt"
	"strconv"
)

type Clock struct {
	min  int
	hour int
}

func (c Clock) hourAsString() string {
	switch {
	case c.hour == 0:
		return "12"
	case c.hour < 10:
		return "0" + strconv.Itoa(c.hour)
	default:
		return strconv.Itoa(c.hour)
	}
}

func (c Clock) minAsString() string {
	switch {
	case c.min < 10:
		return "0" + strconv.Itoa(c.min)
	default:
		return strconv.Itoa(c.min)
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%v:%v", c.hourAsString(), c.minAsString())
}

func (c *Clock) AddMinutes(min int) {
	minutes, hours := c.calculate(min)

	c.hour += hours
	if c.hour > 12 {
		c.hour -= 12
	}
	c.min = minutes
}

func (c *Clock) calculate(min int) (int, int) {
	m := c.min + min
	const hourInMinutes = 60
	h := m / hourInMinutes
	m = m - h*hourInMinutes
	return m, h
}

func NewClock(min, hour int) Clock {
	return Clock{
		min:  min,
		hour: hour,
	}
}
