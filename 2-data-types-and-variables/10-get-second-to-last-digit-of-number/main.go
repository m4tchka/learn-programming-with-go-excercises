package main

import "fmt"

/*
Write a program which gets the second to last digit of a number. In example, for the number 123, the second to last digit is 2.

Hint: Use a combination of / and % operators

Example:
(Given 123)
2
*/

func main() {
	num1 := 123
	secondDigit := (num1 / 10) % 10
	fmt.Println(secondDigit)
}
