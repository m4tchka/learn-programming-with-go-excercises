package main

import "fmt"

// Create a program, which prints the numbers from 1 to 10.

func main() {
	fmt.Println(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// Simple
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// Standard For loop
	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}
	// While loop
}
