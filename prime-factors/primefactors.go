package prime

//TestVersion is the version of the unit tests this will pass
const TestVersion = 1

/*Factors finds all prime factors of a number.*/
func Factors(num int64) []int64 {
	factors := []int64{}
	for f := int64(2); f*f <= num; {
		switch {
		case num%f == 0:
			factors = append(factors, f)
			num /= f
		case f == 2:
			f++
		default:
			f += 2
		}
	}
	if num != 1 {
		factors = append(factors, num)
	}
	return factors
}
