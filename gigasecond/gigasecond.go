package gigasecond

import (
    "time"
    "math"
)

// TestVersion the version of tests to run
const TestVersion = 2 

// AddGigasecond adds one gigasecond to the date
func AddGigasecond(today time.Time) time.Time {
    GIGASECOND := time.Duration(math.Pow(10, 9))
    return today.Add(GIGASECOND * time.Second)
}

// Birthday my birthday to compute a gigasecond
var Birthday, _ = time.Parse("2006-01-02", "1941-12-07")

