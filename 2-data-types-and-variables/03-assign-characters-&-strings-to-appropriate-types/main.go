package main

import "fmt"

/*
Create a program, which assigns lines of text to the appropriate type in Go. Print the line on the terminal. Use either string or character

NOTE: For this exercise, use assignment of the type
var num <some_type> = <some_value>

The numbers:
The gopher is amazing!
G
o
i
s
great
!
*/

func main() {
	var text1 string = "The gopher is amazing!"
	var text2 rune = 'G'
	var text3 rune = 'o'
	var text4 rune = 'i'
	var text5 rune = 's'
	var text6 string = "great"
	var text7 rune = '!'
	fmt.Println(text1)
	fmt.Println(text2, text3, text4, text5, text6, text7)
}
