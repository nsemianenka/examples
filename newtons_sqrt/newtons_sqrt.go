package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	znext, zprev, eps := 1., 2., 0.00001
	for math.Abs(znext-zprev) > eps {
		zprev = znext
		znext = zprev - (zprev*zprev-x)/2/zprev

		//fmt.Printf("znext: %T, %f, zprev: %T, %f, eps: %T, %f\n",
		//		   	znext, znext, zprev, zprev, eps, eps)
	}
	return znext
}

func main() {
	mySqrt, mathSqrt := Sqrt(2), math.Sqrt(2)
	fmt.Printf("my: %F, math: %F, delta: %F\n", mySqrt, mathSqrt,
		math.Abs(mySqrt-mathSqrt))
}
