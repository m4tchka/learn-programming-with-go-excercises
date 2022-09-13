package main

import (
	"fmt"
	"math"
)

/*
Write a program which calculates the third side of a right triangle using the Pythagorean Theorem.

Example:
(Given 3 4)
5

(Given 6 7)
9.219544457292887
*/

func main() {
	sideA := 3
	sideB := 4
	sideC := math.Sqrt((float64(sideA)*float64(sideA) + float64(sideB)*float64(sideB)))
	fmt.Printf("A: %v, B: %v, C: %v\n", sideA, sideB, sideC)

	sideD := 6
	sideE := 7
	sideF := math.Sqrt(math.Pow(float64(sideD), 2) + math.Pow(float64(sideE), 2))
	fmt.Printf("D: %v, E: %v, F: %v\n", sideD, sideE, sideF)
}
