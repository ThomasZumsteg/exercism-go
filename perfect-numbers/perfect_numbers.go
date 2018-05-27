package perfect

import "errors"

type Classification string

var ErrOnlyPositive error = errors.New("Must be 1 or greater")

const (
	ClassificationPerfect   Classification = "perfect"
	ClassificationAbundant  Classification = "abundant"
	ClassificationDeficient Classification = "deficient"
)

func Classify(number int64) (result Classification, err error) {
	if number <= 0 {
		err = ErrOnlyPositive
		return
	} else if number == 1 {
		result = ClassificationDeficient
		return
	}

	var total int64 = 1
	for i := int64(2); i*i < number; i++ {
		if number%i == 0 {
			total += i
			if i*i != number {
				total += number / i
			}
		}
	}
	if total < number {
		result = ClassificationDeficient
	} else if number < total {
		result = ClassificationAbundant
	} else if number == total {
		result = ClassificationPerfect
	}
	return
}
