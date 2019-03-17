package learntest

import (
	"errors"
	"math"
)

//Plusone adds one to the integer
func Plusone(w int) int {
	z := w + 1
	return z
}

// SquareRoot accepts int and returns float64 squareroot
func SquareRoot(w int) (float64, error) {
	if w < 0 {
		return 0, errors.New("Input smaller than zero")
	}
	x := float64(w)
	return math.Sqrt(x), nil
}
