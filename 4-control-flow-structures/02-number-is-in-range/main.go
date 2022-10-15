package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which checks if a given input number x is in the range [a; b). That is a inclusively and b exclusively.
x, a and b are given as input in the terminal.
*/

func main() {
	x := getInput()
	a := getInput()
	b := getInput()
	if x >= a && x < b {
		fmt.Printf("\n%v is in the range [%v; %v)", x, a, b)
	} else {
		fmt.Printf("\n%v is not in the range [%v; %v)", x, a, b)
	}
}
func getInput() int {
	r := bufio.NewReader(os.Stdin)
	numStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(numStr))
	if err != nil {
		panic(err)
	}
	return num
}
