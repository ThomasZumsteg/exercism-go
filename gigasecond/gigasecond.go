package gigasecond

import (
	"math"
	"time"
)

// TestVersion is the version of the unit tests was designed to pass.
const TestVersion = 2

/*AddGigasecond adds one gigasecond (1E9 seconds) to the date.*/
func AddGigasecond(today time.Time) time.Time {
	GIGASECOND := time.Duration(math.Pow(10, 9))
	return today.Add(GIGASECOND * time.Second)
}

// Birthday is my birthday and is used to compute a giga-versery.
var Birthday, _ = time.Parse("2006-01-02", "1941-12-07")
