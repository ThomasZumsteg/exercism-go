package armstrong

func IsNumber(number int) bool {
	digits := []int{}
	remainer := number
	numDigits := 0
	for ; remainer > 0; numDigits++ {
		digits = append(digits, remainer%10)
		remainer = remainer / 10
	}

	armstrongSum := 0
	for _, d := range digits {
		armstrongSum += powInt(d, numDigits)
	}

	return number == armstrongSum
}

func powInt(num int, pow int) int {
	result := 1
	for i := 0; i < pow; i++ {
		result *= num
	}
	return result
}
