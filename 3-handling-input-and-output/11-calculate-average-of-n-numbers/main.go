package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Create a program which reads in N numbers from the terminal, given on a single line and calculates their average.
*/

func main() {
	average := getInput()
	fmt.Printf("Average = %.2f", average)
}

func getInput() float64 {
	r := bufio.NewReader(os.Stdin)
	numsStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	numsStrSlice := strings.Split(strings.TrimSpace(numsStr), " ")
	var sum float64 = 0
	for _, numStr := range numsStrSlice {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	// fmt.Printf("Sum: %.2f\nLen: %d\n", sum, len(numsStrSlice))
	return sum / float64(len(numsStrSlice))
}
