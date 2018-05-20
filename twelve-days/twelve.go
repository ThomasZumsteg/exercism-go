package twelve

import (
	"fmt"
	"strings"
)

var days = [][]string{
	[]string{"first", "a Partridge in a Pear Tree."},
	[]string{"second", "two Turtle Doves"},
	[]string{"third", "three French Hens"},
	[]string{"fourth", "four Calling Birds"},
	[]string{"fifth", "five Gold Rings"},
	[]string{"sixth", "six Geese-a-Laying"},
	[]string{"seventh", "seven Swans-a-Swimming"},
	[]string{"eighth", "eight Maids-a-Milking"},
	[]string{"ninth", "nine Ladies Dancing"},
	[]string{"tenth", "ten Lords-a-Leaping"},
	[]string{"eleventh", "eleven Pipers Piping"},
	[]string{"twelfth", "twelve Drummers Drumming"},
}

const first_line = "On the %s day of Christmas my true love gave to me"

func Verse(verse int) string {
	lines := []string{fmt.Sprintf(first_line, days[verse-1][0])}
	for d := verse - 1; 0 <= d; d-- {
		if d == 0 && verse != 1 {
			lines = append(lines, "and "+days[d][1])
		} else {
			lines = append(lines, days[d][1])
		}
	}
	return strings.Join(lines, ", ")
}

func Song() string {
	var song []string
	for i := 1; i <= 12; i++ {
		song = append(song, Verse(i)+"\n")
	}
	return strings.Join(song, "")
}
