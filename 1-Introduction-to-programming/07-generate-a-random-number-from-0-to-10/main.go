package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Create a go program which generates a random number between 0 and 10 inclusively.
//Print the number on the terminal
//Every time you run the program, a different number should be shown.

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(11))
}
