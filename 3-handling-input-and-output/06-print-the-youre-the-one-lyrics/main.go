package main

import "fmt"

/*
Create a program which initializes a string variable with the contents “the one”.

Use it to print the following lines from the famous “The Vogues” song:
You’re {str} that I long to kiss, Baby you’re {str} that I really miss. You’re {str} that I’m dreamin’ of, Baby, you’re {str} that I love!
*/

func main() {
	str := "the one"
	fmt.Printf("You’re %s that I long to kiss, Baby you’re %s that I really miss. You’re %s that I’m dreamin’ of, Baby, you’re %s that I love!", str, str, str, str)
}
