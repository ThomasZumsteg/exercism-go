package hexadecimal

//rangeErrors are when a hex number is too big
type rangeError struct {
	message string
}

/*Error makes a rangeError into a string.*/
func (err rangeError) Error() string {
	return err.message
}

/*syntaxErrors are when a hex number isn't properly formatted.*/
type syntaxError struct {
	message string
}

/*Error make syntaxErrors into a string.*/
func (err syntaxError) Error() string {
	return err.message
}

/*ParseHex converst a hexidecimal string to an integer.*/
func ParseHex(hexString string) (int64, error) {
	if hexString == "" {
		return 0, syntaxError{"No digits"}
	} else if 16 <= len(hexString) {
		return 0, rangeError{"Number is too big, can only handle 16 digits"}
	}
	var dec int64
	for _, hexDigit := range hexString {
		dec <<= 4
		switch hexDigit {
		case '0':
			dec |= 0
		case '1':
			dec |= 1
		case '2':
			dec |= 2
		case '3':
			dec |= 3
		case '4':
			dec |= 4
		case '5':
			dec |= 5
		case '6':
			dec |= 6
		case '7':
			dec |= 7
		case '8':
			dec |= 8
		case '9':
			dec |= 9
		case 'A', 'a':
			dec |= 10
		case 'B', 'b':
			dec |= 11
		case 'C', 'c':
			dec |= 12
		case 'D', 'd':
			dec |= 13
		case 'E', 'e':
			dec |= 14
		case 'F', 'f':
			dec |= 15
		default:
			return 0, syntaxError{"Not a digit"}
		}
	}
	return dec, nil
}

/*HandleErrors report errors converting hex numbers.*/
func HandleErrors(hexStrings []string) []string {
	errorList := make([]string, len(hexStrings))
	for i, hexString := range hexStrings {
		_, err := ParseHex(hexString)
		switch err.(type) {
		case rangeError:
			errorList[i] = "range"
		case syntaxError:
			errorList[i] = "syntax"
		default:
			errorList[i] = "none"
		}
	}
	return errorList
}
