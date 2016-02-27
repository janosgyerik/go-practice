package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

const epsilon = 1e-15

func isGood(x float64, p float64) bool {
	return math.Abs(p * p - x) < epsilon
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	p := x / 2;
	for ; !isGood(x, p); {
		p = (x / p + p) / 2;
	}

	return p, nil
}

func main() {
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
