package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Create a function incrementBy which takes in a slice and a number N and increments all the elements of the slice by N.
*/

func main() {
	sli, n := getIntSliInput(), getNumInput()
	newSli := incrementBy(sli, n)
	fmt.Println("result: ", newSli)
}
func incrementBy(sli []int, N int) []int {
	for ind := range sli {
		sli[ind] += N
	}
	return sli
}
func getIntSliInput() []int {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	var intSli []int
	strSli := strings.Split(strings.TrimSpace(line), " ")
	for _, str := range strSli {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intSli = append(intSli, num)
	}
	return intSli
}
func getNumInput() int {
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
