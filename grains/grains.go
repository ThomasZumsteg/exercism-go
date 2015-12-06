package grains

import (
	"errors"
	"math"
)

/*Square computes the number of grains on a square of a chessboard
where the number on each square doubles.
Only valid for the number of squres on a chessboard 1-64*/
func Square(square int) (uint64, error) {
	if !(1 <= square && square <= 64) {
		return 0, errors.New("Not a valid square")
	}
	return uint64(math.Pow(2, float64(square-1))), nil
}

/*Total computes to total number of grains on a chessboard
where the number of grains on each successive square doubles*/
func Total() uint64 {
	return uint64(math.Pow(2, float64(65))) - 1
}
