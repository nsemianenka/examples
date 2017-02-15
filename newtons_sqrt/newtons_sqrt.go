package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e));
}

func Sqrt(x float64) (float64, error) {
	znext, zprev, eps := 1., 2., 0.00001
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for math.Abs(znext-zprev) > eps {
		zprev = znext
		znext = zprev - (zprev*zprev-x)/2/zprev
	}

	return znext, nil
}

func main() {
	mySqrt, err := Sqrt(2)
	mathSqrt := math.Sqrt(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("my: %F, math: %F, delta: %F\n", mySqrt, mathSqrt,
			math.Abs(mySqrt-mathSqrt))
	}
}
