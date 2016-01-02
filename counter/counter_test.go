package counter

import "testing"

var tests = []struct {
	input                      string
	lines, letters, characters int
}{
	{"simple test\n", 1, 10, 12},
	{"two\nlines", 2, 8, 9},
	{"ends with \\n\n", 1, 9, 13},
	{"üñîćøđę\n", 1, 7, 8},
	{"", 0, 0, 0},
}

func TestLines(t *testing.T) {
	for _, test := range tests {
		counter := makeCounter()
		counter.AddString(test.input)
		if actual := counter.Lines(); actual != test.lines {
			t.Errorf("counter.Lines(%q) = %d, expected %d.",
				test.input, actual, test.lines)
		}
	}
}

func TestLetters(t *testing.T) {
	for _, test := range tests {
		counter := makeCounter()
		counter.AddString(test.input)
		if actual := counter.Letters(); actual != test.letters {
			t.Errorf("counter.Letters(%q) = %d, expected %d.",
				test.input, actual, test.letters)
		}
	}
}

func TestCharacters(t *testing.T) {
	for _, test := range tests {
		counter := makeCounter()
		counter.AddString(test.input)
		if actual := counter.Characters(); actual != test.characters {
			t.Errorf("counter.Characters(%q) = %d, expected %d.",
				test.input, actual, test.characters)
		}
	}
}
