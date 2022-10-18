package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which takes as input a number N and generates all numbers in the range [1; N] which are prime numbers.
Outputting them on a single line is optional, I’ve written it that way just to save some space.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	for i := 2; i < num; i++ {
		var isIPrime = false
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isIPrime = true
			}
		}
		if !isIPrime {
			fmt.Printf("%d ", i)
		}
	}
}
