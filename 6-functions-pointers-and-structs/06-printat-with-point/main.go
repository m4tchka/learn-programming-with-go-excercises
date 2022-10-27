package main

import (
	"time"

	tm "github.com/buger/goterm"
)

/*
Modify the PrintAt function you created in Problem 02 to accept a Point instead of a row and col. The function should still work as expected.
*/

func main() {
	tm.Clear()
	printAt(Point{5, 5}, "(")
	printAt(Point{5, 6}, "*")
	printAt(Point{5, 7}, ")")
	time.Sleep(time.Second)
}
func printAt(pt Point, str string) {
	tm.MoveCursor(pt.col, pt.row)
	tm.Print(str)
	tm.Flush()
}

type Point struct {
	row int
	col int
}
