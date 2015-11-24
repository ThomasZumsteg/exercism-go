package leap

// TestVersion Version of the test suite that the program was build against
const TestVersion = 1

/*
IsLeapYear
Returns true if the "year" is a leap year, false otherwise

Years are determined by the following rules
The year is evenly divisble by 4 and not by 100
unless it's evenly divisible by 400
*/
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
