package meetup

import (
	"fmt"
	"time"
)

//TestVersion is the unit test verions that this will work for.
const TestVersion = 1

//WeekSchedule is the position of a week in the month
type WeekSchedule int

//The valid list of positions
const (
	First  WeekSchedule = iota // 1-7
	Second                     // 8-14
	Third                      // 15-21
	Fourth                     // 22-29
	Last                       // [31|30|29|28] - [25|24|23|22]
	Teenth                     // 13-19
)

/*MeetupDay finds a day in a month that is in some position in the month.*/
func MeetupDay(nth WeekSchedule, day time.Weekday, month time.Month, year int) int {
	strStart := fmt.Sprintf("%02d-%02d-%04d", 1, month, year)
	start, parseErr := time.Parse("02-01-2006", strStart)
	if parseErr != nil {
		return 0
	}

	switch nth {
	case First:
		return dayInWeek(day, start.AddDate(0, 0, 0), false).Day()
	case Second:
		return dayInWeek(day, start.AddDate(0, 0, 7), false).Day()
	case Third:
		return dayInWeek(day, start.AddDate(0, 0, 14), false).Day()
	case Fourth:
		return dayInWeek(day, start.AddDate(0, 0, 21), false).Day()
	case Teenth:
		return dayInWeek(day, start.AddDate(0, 0, 12), false).Day()
	case Last:
		return dayInWeek(day, start.AddDate(0, 1, -1), true).Day()
	default:
		return 0
	}
}

/*dayInWeek search a week for a particular weekday.*/
func dayInWeek(day time.Weekday, start time.Time, reverse bool) time.Time {
	step := 1
	if reverse {
		step = -1
	}
	for d := 0; d < 7; d++ {
		date := start.AddDate(0, 0, d*step)
		if date.Weekday() == day {
			return date
		}
	}
	return start
}
