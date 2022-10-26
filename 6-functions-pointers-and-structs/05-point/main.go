package main

import "fmt"

/*
Create a struct called Point which takes in two parameters - row and col.
Create several instances of Point and print them on the terminal.
*/
func main() {
	p1 := Point{row: 1, col: 2}
	p2 := Point{row: 2, col: 4}
	p3 := Point{row: 10, col: 30}
	p4 := Point{col: 10, row: 30}
	var ptSli []Point
	ptSli = append(ptSli, p1, p2, p3, p4)

	for _, pt := range ptSli {
		fmt.Println(pt)
	}
}

type Point struct {
	row int
	col int
}
