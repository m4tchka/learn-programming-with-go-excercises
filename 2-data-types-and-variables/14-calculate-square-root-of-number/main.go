package main

import (
	"fmt"
	"math"
)

/*
Write a program which calculates the square root of a number.

Hint: Don’t try implementing your own “square root”. Try googling around for an already provided solution

Example:
(Given 9)
3
(Given 13)
3.605551275463989
*/

func main() {
	num1 := 9
	num2 := 13
	sqrt1 := math.Sqrt(float64(num1))
	sqrt2 := math.Sqrt(float64(num2))
	fmt.Println(sqrt1, sqrt2)
}
