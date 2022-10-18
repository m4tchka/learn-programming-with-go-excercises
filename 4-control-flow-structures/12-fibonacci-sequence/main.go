package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which takes as input a number N and generates the fibonacci sequence up to N.
In the fibonacci sequence, the nth number is equal to the sum of the previous two numbers.
The first and second numbers are always equal to 1.
*/

/*
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
		var res []int
		res = append(res, 1)
		res = append(res, 1)
		for i := 2; i <= num; i++ {
			res = append(res, res[i-2]+res[i-1])
		}
		fmt.Sprintf(res)
	}
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
	first, second := 1, 1
	fmt.Printf("%d %d", first, second)
	for i := 2; i <= num; i++ {
		next := first + second
		fmt.Printf(" %d", next)
		first = second
		second = next
	}
}
