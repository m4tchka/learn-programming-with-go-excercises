package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which checks if a given input number x is in a given range. The bounds of the range are given as the flags in/ex which specify if the ranges are inclusive or not.
The output should be formatted as {x} is (not) in the range {range} where range has different brackets based on the in/ex flags.
For example, if the flags given are in + ex, then the range would be [a; b).
If they are in in, the range would be [a; b].
*/

func main() {
	x := getNumInput()
	a := getNumInput()
	b := getNumInput()
	isStartInclusive := getIncExc()
	isEndInclusive := getIncExc()
	fmt.Printf("\nx: %d, a: %d, b: %d", x, a, b)
	fmt.Printf("\niSI: %v, iEI: %v", isStartInclusive, isEndInclusive)
	var startCond, endCond bool
	if isStartInclusive {
		startCond = x >= a
	} else {
		startCond = x > a
	}
	if isEndInclusive {
		endCond = x <= b
	} else {
		endCond = x < b
	}
	var startChar, endChar rune
	if isStartInclusive {
		startChar = '['
	} else {
		startChar = '('
	}
	if isEndInclusive {
		endChar = ']'
	} else {
		endChar = ')'
	}
	if startCond && endCond {
		fmt.Printf("\n%d is in the range %c%d; %d%c", x, startChar, a, b, endChar)
	} else {
		fmt.Printf("\n%d is not in the range %c%d; %d%c", x, startChar, a, b, endChar)
	}
}
func getNumInput() int {
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
func getIncExc() bool {
	r := bufio.NewReader(os.Stdin)
	incOrExc, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	isRangeInc := strings.TrimSpace(incOrExc) == "in"
	return isRangeInc
}
