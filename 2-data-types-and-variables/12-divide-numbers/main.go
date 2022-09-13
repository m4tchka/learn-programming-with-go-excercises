package main

import "fmt"

/*
Create a program which divides two numbers as integers and then as floating-point nums.

Example:
(Given 13 3)
4
4.333333333333333
*/

func main() {
	var num1 int = 13
	var num2 int = 3
	result := num1 / num2
	fmt.Printf("%v divided by %v gives %v of type %T\n", num1, num2, result, result)

	var num3 float64 = 13
	var num4 float64 = 3
	result2 := num3 / num4
	fmt.Printf("%v divided by %v gives %v of type %T\n", num3, num4, result2, result2)
}
