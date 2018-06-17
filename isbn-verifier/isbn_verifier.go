package isbn

func IsValidISBN(isbn string) bool {
	digits, checksum := 0, 0
	for _, digit := range isbn {
		if '0' <= digit && digit <= '9' {
			checksum += int(digit-'0') * (10 - digits)
			digits++
		} else if 'X' == digit {
			if digits != 9 {
				return false
			}
			checksum += 10 * (10 - digits)
			digits++
		}
	}
	return digits == 10 && checksum%11 == 0
}
