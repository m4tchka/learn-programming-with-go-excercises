package main

import "fmt"

/*
Create a program, which assigns a set of real number values to the appropriate type in Go. Print the number on the terminal. The numbers should be stored in the smallest type possible.

NOTE: Don’t simply use strings everywhere
NOTE: For this exercise, use assignment of the type
var num <some_type> = <some_value>

The numbers:
14.435234
14.435234234234235
*/

func main() {
	var num1 float32 = 14.435234
	var num2 float64 = 14.435234234234235
	fmt.Println("num1: ", num1, "num2: ", num2)
}
