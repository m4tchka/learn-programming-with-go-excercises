package main

import (
	"fmt"
	"math"
)

/*
Write a program which compares two float numbers with a given acceptable delta of 0.00005.

In example, the numbers 3.100005 and 3.100003 are considered equal as their difference is less than the given delta.

Hint: for this task, you’ll need to learn about if-else statements.

Examples:
3.100005 is equal to 3.100003
1.12003 is not equal to 1.12009
*/

func main() {
	n1 := 3.100005
	n2 := 3.100003
	d := 0.000005
	if math.Abs(n1-n2) <= d {
		fmt.Printf("%v is equal to %v\n", n1, n2)
	} else {
		fmt.Printf("%v is not equal to %v\n", n1, n2)
	}
}
