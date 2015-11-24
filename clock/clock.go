package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	minutes int
}

func Time(hour, minute int) Clock {
	time := (hour*60 + minute) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock{time}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

func (c Clock) Add(minutes int) Clock {
	time := (c.minutes + minutes) % (60 * 24)
	if time < 0 {
		time += 60 * 24
	}
	return Clock{time}
}
