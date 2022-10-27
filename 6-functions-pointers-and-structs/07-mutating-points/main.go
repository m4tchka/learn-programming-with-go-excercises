package main

import "fmt"

/*
Create three functions - incPtr, incNoPtr and incNewPtr which all take in a point and increment its values by 1. Here’s the catch:
incPtr should take in a pointer to a point & modify its contents
incNoPtr should take a point without a pointer & modify its contents
incNewPtr should take in a pointer and assign it a new Point with the incremented values
Print the values of the three points after being incremented. Anything unusual?
*/

func main() {
	pt1 := Point{col: 1, row: 1}
	pt2 := Point{col: 1, row: 1}
	pt3 := Point{col: 1, row: 1}
	incPtr(&pt1)
	incNoPtr(pt2)
	incNewPtr(&pt3)
	fmt.Println(pt1, pt2, pt3)
}
func incPtr(pt *Point) {
	pt.col++
	pt.row++
}
func incNoPtr(pt Point) {
	pt.col++
	pt.row++
}
func incNewPtr(pt *Point) { // Pass a pointer to a Point as the argument of the function. This variable is a reference to the original value.
	ptx := &Point{col: 2, row: 2} // Create a new pointer that references a Point with new values
	pt = ptx                      // Make the argument to reference the new pointer (WITHIN THE FUNCTION SCOPE ONLY)
	pt.row++                      // Increment the Point's fields
	pt.col++
}

type Point struct {
	col int
	row int
}
