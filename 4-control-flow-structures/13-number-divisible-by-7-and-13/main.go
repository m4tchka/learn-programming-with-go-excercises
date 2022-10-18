package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which takes as input a number N and outputs all numbers in the range [1; N] which are divisible by both 7 and 13.
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
	for i := 1; i <= num; i++ {
		if i%7 == 0 && i%13 == 0 {
			fmt.Println(i)
		}
	}
}
