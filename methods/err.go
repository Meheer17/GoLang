package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return float64(0), ErrNegativeSqrt(x)
	}
	
	z := x / 2
    for i := 0; i < 100; i++ {
        z -= (z*z - x) / (2 * z)
    }
	return z, nil 
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
