package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Read three numbers from the terminal & calculate their sum.
The numbers should be input one at a line.

Afterwards, read three more number, input on a single line and calculate their sum.
*/

func main() {
	num1 := part1Input()
	num2 := part1Input()
	num3 := part1Input()
	part1Sum := num1 + num2 + num3
	fmt.Printf("Result: %d\n", part1Sum)
	part2Sum := part2Input()
	fmt.Printf("Result: %d\n", part2Sum)

}
func part1Input() int {
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
func part2Input() int {
	r := bufio.NewReader(os.Stdin)
	inputStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	numbers := strings.Split(strings.TrimSpace(inputStr), " ")
	var sum int
	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	return sum
}
