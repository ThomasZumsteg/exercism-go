package leap

const TestVersion = 1

// True if the year is a leap year
// leap years are evenly divisible by 4
// unless they are evenly divisible by 100
// and not evenly divisible by 400
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
