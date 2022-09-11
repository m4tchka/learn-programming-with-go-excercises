package main

import "fmt"

/*
Create a program, which assigns a set of integer values to the appropriate type in Go. Print the integer on the terminal. The numbers should be stored in the smallest type possible.
NOTE: Don’t simply use strings everywhere
NOTE: For this exercise, use assignment of the type
var num <some_type> = <some_value>

The numbers:
60
-100
127
128
-144243
255
256
144243
3641
-4512
65535
10000000000000000000
65536
-1000000000000000000
-30000
*/

func main() {
	var num1 int8 = 60
	var num2 int16 = -100
	var num3 int8 = 127
	var num4 uint8 = 128
	var num5 int32 = -144243
	var num6 uint8 = 255
	var num7 uint16 = 256
	var num8 uint32 = 144243
	var num9 int16 = 3641
	var num10 int16 = -4512
	var num11 uint16 = 65535
	var num12 uint64 = 10000000000000000000
	var num13 uint32 = 65536
	var num14 int64 = -1000000000000000000
	var num15 int16 = -30000
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)
	fmt.Println(num8)
	fmt.Println(num9)
	fmt.Println(num10)
	fmt.Println(num11)
	fmt.Println(num12)
	fmt.Println(num13)
	fmt.Println(num14)
	fmt.Println(num15)
}
