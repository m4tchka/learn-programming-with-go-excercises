package main

import "fmt"

/*
Write a program which converts from kg to lbs. 1 kg is equal to 2.2046226218 lbs.

Example:
(Given 1)
2.2046226218

(Given 78.5)
173.0628758113
*/

func main() {
	const kgLbsRatio = 2.2046226218
	var num1 float32 = 1
	lbs1 := num1 * kgLbsRatio
	fmt.Println(lbs1)
	var num2 float64 = 78.5
	lbs2 := num2 * kgLbsRatio
	fmt.Println(lbs2)
}
