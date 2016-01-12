package ocr

import "strings"

/*recognizeDigit checks if a pipes and bars character is a digit.*/
func recognizeDigit(digit string) (string, bool) {
	d, ok := digits[digit]
	return d, ok
}

/*Recognize reads pipes and bars characters from a string.
Characters are 3 wide and 4 high.*/
func Recognize(characters string) []string {
	var digits []string
	for l, line := range groupLines(characters[1:]) {
		digits = append(digits, "")
		for _, digit := range readLine(line) {
			if d, ok := recognizeDigit(digit); ok {
				digits[l] += d
			} else {
				digits[l] += "?"
			}
		}
	}
	return digits
}

/*groupLines breaks a string into groups of lines 4 lines long.*/
func groupLines(str string) [][]string {
	var groups [][]string
	for l, line := range strings.Split(str, "\n") {
		if l%4 == 0 {
			groups = append(groups, []string{})
		}
		groups[l/4] = append(groups[l/4], line)
	}
	return groups
}

/*readLine reads characters from a list of lines.*/
func readLine(lines []string) []string {
	var chars []string
	for c := 0; c+3 <= len(lines[0]); c += 3 {
		char := ""
		for l := 0; l < 4; l++ {
			char += "\n" + lines[l][c:c+3]
		}
		chars = append(chars, char)
	}
	return chars
}

//digits maps OCR digits to their string equivelent
var digits = map[string]string{
	zero: "0", one: "1", two: "2", three: "3", four: "4",
	five: "5", six: "6", seven: "7", eight: "8", nine: "9",
}

//zero, one, two, three, etc. define the OCR digits
const zero string = `
 _ 
| |
|_|
   `
const one string = `
   
  |
  |
   `
const two string = `
 _ 
 _|
|_ 
   `
const three string = `
 _ 
 _|
 _|
   `
const four string = `
   
|_|
  |
   `
const five string = `
 _ 
|_ 
 _|
   `
const six string = `
 _ 
|_ 
|_|
   `
const seven string = `
 _ 
  |
  |
   `
const eight string = `
 _ 
|_|
|_|
   `
const nine string = `
 _ 
|_|
 _|
   `
