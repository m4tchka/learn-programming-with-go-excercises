package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Read three real numbers from the terminal & calculate their sum.
The output should be formatted with three decimal places.
*/
func main() {
	num1 := getInput()
	num2 := getInput()
	num3 := getInput()
	avg := (num1 + num2 + num3) / 3
	fmt.Printf("Result: %.3f", avg)
}

func getInput() float64 {
	r := bufio.NewReader(os.Stdin)
	numStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	numFlt, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
	if err != nil {
		panic(err)
	}
	return numFlt
}
