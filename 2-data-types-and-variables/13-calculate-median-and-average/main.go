package main

import "fmt"

/*
Write a program which calculates the median and average of three numbers.

Example:
(Given 3 18 23)
18
14.666666666666666
*/
func main() {
	num1 := 3
	num2 := 18
	num3 := 23
	median := num2
	avg := (num1 + num2 + num3) / 3
	fmt.Println(median, avg)
}
