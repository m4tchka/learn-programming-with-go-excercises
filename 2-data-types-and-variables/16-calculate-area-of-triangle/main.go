package main

import "fmt"

/*
Write a program which calculates the area of a triangle, given base and height.

Example:
(Given 3 5)
7.5
*/

func main() {
	num1 := 3
	num2 := 5
	var area float64 = (float64(num1) * float64(num2)) / 2
	fmt.Println(area)
}
