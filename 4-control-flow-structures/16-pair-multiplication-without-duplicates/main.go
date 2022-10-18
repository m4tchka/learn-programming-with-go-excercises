package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Extend the previous solution to not allow duplicate values.
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
	for i := 1; i <= int(math.Sqrt(float64(N))); i++ {
		for j := N; j >= int(math.Sqrt(float64(N))); j-- {
			if i*j == X {
				fmt.Printf("\n%d x %d = %d", i, j, X)
			}
		}
	}
}
