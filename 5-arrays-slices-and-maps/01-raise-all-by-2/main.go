package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which reads a set of integers, on a single line, stores the numbers in a slice & raises all elements’ power by 2.
You are given N which represents the count of numbers which come after it.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	numStrSli := strings.Split(strings.TrimSpace(line), " ")
	if err != nil {
		panic(err)
	}
	var numSli []int
	for i := range numStrSli {
		num, err := strconv.Atoi(numStrSli[i])
		if err != nil {
			panic(err)
		}
		numSli = append(numSli, num)
	}
	for _, num := range numSli {
		num *= num
		fmt.Printf("%d ", num)
	}
}
