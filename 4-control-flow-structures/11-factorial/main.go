package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which takes as input a number N and calculates N! (factorial).
The factorial of a number is the multiplication of the numbers [1; N]
*/

/* func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	var result int = 1
	for i := num; i > 0; i-- {
		result *= i
	}
	fmt.Println(result)
} */

func main() {
	n := getInput()
	f := calculateFactorial(n, 1)
	fmt.Printf("\nFactorial: %d", f)
}
func getInput() int {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	return num
}
func calculateFactorial(num, total int) int {
	if num == 1 {
		return total
	} else {
		return calculateFactorial(num-1, total*num)
	}
}
