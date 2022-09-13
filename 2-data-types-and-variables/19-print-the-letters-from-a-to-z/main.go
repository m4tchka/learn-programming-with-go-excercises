package main

import (
	"fmt"
)

/*
Write a program which prints all the letters in the english alphabet.
To do this, learn about the ASCII table.

Use a for loop to iterate through all characters from a to z and print them out on separate lines.

Hint: ‘a’ is represented in memory as an integer. You can use it to initialize your loop variable & go all the way to ‘z’

Example:
a
b
c
d
(remainder omitted...)
x
y
z
*/

/*
	 func main() {
		for i := 97; i <= 122; i++ {
			fmt.Println(string(rune(i)))
		}
	}
*/
func main() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Println(string(i))
	}
}
