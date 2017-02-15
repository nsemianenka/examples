package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e);
}

func Sqrt(x float64) (float64, error) {
	var error *ErrNegativeSqrt
	znext, zprev, eps := 1., 2., 0.00001
	if x < 0 {
		error = &ErrNegativeSqrt(x)
	}
	if error == nil {
		for math.Abs(znext-zprev) > eps {
			zprev = znext
			znext = zprev - (zprev*zprev-x)/2/zprev

			//fmt.Printf("znext: %T, %f, zprev: %T, %f, eps: %T, %f\n",
			//		   	znext, znext, zprev, zprev, eps, eps)
		}
	}
	return znext, error
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
