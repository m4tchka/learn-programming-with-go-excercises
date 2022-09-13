package main

import (
	"fmt"
	"math"
)

/*
Write a program which checks whether a point is inside a circle. You are given a point P and the center of a circle C at (0, 0) with a radius 4.

See all the examples in this desmos graph.

Examples:
(Given (-2,1))
true

(Given (1, 4))
false

(Given (2,2))
true

(Given (-2,-2))
true
*/
func main() {
	//x^2 + y^2 <= radius^2
	//sqrt(x^2 + y^2)<=radius
	var x1 float64 = -2
	var x2 float64 = 1
	var rad float64 = 4
	fmt.Println(math.Sqrt(math.Pow(x1, 2)+math.Pow(x2, 2)) <= rad)
}
