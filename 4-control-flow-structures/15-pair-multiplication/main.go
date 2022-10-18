package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which takes as input two numbers - N and X and find all pairs in the range [1; N], which when multiplied, produce X.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	N, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	line2, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	X, err := strconv.Atoi(strings.TrimSpace(line2))
	if err != nil {
		panic(err)
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if i*j == X {
				fmt.Printf("\n%d x %d = %d", i, j, X)
			}
		}
	}
}
