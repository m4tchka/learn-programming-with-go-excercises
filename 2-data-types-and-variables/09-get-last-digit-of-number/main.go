package main

import "fmt"

/*
Create a program which gets the last digit of a number. In example, for the number 123, the last digit is 3.

Hint: Use the % operator

Example:
(Given 123)
3
*/
func main() {
	var num int = 123
	var lastDigit int = num % 10
	fmt.Println(lastDigit)
}
