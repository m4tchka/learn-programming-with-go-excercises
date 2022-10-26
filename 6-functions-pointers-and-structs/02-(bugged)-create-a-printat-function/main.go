package main

import (
	"time"

	tm "github.com/buger/goterm"
	"github.com/fatih/color"
)

/*
Using the buger/goterm library, create a PrintAt function which takes in a row, col and string and prints the string at the designated position on the terminal.
*/

//	func main() {
//		PrintAt(5, 5, "Hello")
//	}
//
//	func PrintAt(row int, col int, str string) {
//		fmt.Printf("(invoked PrintAt(%d, %d, %s))\n", row, col, str)
//		tm.Clear()
//		tm.MoveCursor(1, 1)
//		tm.Print()
//	}
func main() {
	color.Red("Test")
	tm.Clear() // Clear current screen

	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 15)

		tm.Println("Current Time:", time.Now().Format(time.RFC1123))

		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(time.Second)
	}
}
