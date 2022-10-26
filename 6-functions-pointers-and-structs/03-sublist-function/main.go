package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Create a program which takes in a slice as input and returns a  sub-slice as output. It should accept two parameters, other than the slice - start and end.
The part of the slice which should be returned is [start; end).
Print the returned subslice on the terminal.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	sli := strings.Split(strings.TrimSpace(line), " ")
	s, e := getNumInput(), getNumInput()
	subSli := subslice(sli, s, e)
	fmt.Println(strings.Join(subSli, ", "))
}
func subslice(s []string, start int, end int) []string {
	fmt.Printf("\nInvoking subslice(%s, %d, %d)\n", fmt.Sprint(s), start, end)
	return s[start:end]
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
