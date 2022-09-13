package main

import "fmt"

/*
Convert the following values to the type specified & print them out:
float32 with value 3.14 -> int
float64 with value 12.3456789 -> float32
int16 with value 1234 -> int8
int16 with value 1234 -> uint8
*/

func main() {
	var num1 float32 = 3.14
	newNum1 := int(num1)
	fmt.Printf("num1: %v, newNum1: %v\n", num1, newNum1)
	var num2 float64 = 12.3456789
	newNum2 := float32(num2)
	fmt.Printf("num2: %v, newNum2: %v\n", num2, newNum2)
	var num3 int16 = 1234
	newNum3 := int8(num3)
	fmt.Printf("num3: %v, newNum3: %v\n", num3, newNum3)
	var num4 int16 = 1234
	newNum4 := uint8(num4)
	fmt.Printf("num4: %v, newNum4: %v\n", num4, newNum4)
}
