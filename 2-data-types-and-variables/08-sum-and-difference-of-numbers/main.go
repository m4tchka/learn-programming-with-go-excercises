package main

import "fmt"

/*
Create a program which calculates the sum and difference of three numbers you choose.

Example:
(Given 5, 7, 8)
20
-10
*/
func main() {
	var num1 int = 5
	num2 := 7
	var num3 int = 8

	sum := num1 + num2 + num3
	diff := num1 - num2 - num3
	fmt.Printf("Sum: %v, Difference: %v", sum, diff)
}
