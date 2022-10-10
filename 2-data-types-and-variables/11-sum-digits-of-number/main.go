package main

import "fmt"

/*
Write a program which sums the digits of a given number. It should work for numbers with arbitrary digit count.

Example:
(Given 1234)
10
*/

func main() {
	sum := 0
	num1 := 123456
	for num1 > 0 {
		lastDigit := num1 % 10
		sum += lastDigit
		num1 /= 10
	}
	fmt.Println(sum)
}
