package main

import (
	"time"

	tm "github.com/buger/goterm"
)

/*
Using the buger/goterm library, create a PrintAt function which takes in a row, col and string and prints the string at the designated position on the terminal.
*/

func main() {
	tm.Clear()
	printAt(5, 5, "(")
	printAt(5, 6, "*")
	printAt(5, 7, ")")
	time.Sleep(time.Second)
}
func printAt(row, col int, str string) {
	tm.MoveCursor(col, row)
	tm.Print(str)
	tm.Flush()
}
