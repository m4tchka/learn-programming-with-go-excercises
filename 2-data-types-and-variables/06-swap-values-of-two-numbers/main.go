package main

import "fmt"

/*
Declare two variables a and b with values 14 and 32. Swap their values such that a contains 32 and b contains 14.

Example:
Before swapping:
a = 14
b = 32

After swapping
a = 32
b = 14

Print the resulting numbers on the terminal.
Extra Challenge: Find a way to do this without an extra variable
*/
func main() {
	a := 14
	b := 32
	fmt.Printf("a: %v, b: %v\n", a, b)

	/* 	c := a
	   	a = b
	   	b = c */
	a, b = b, a
	fmt.Printf("a: %v, b: %v\n", a, b)
}
