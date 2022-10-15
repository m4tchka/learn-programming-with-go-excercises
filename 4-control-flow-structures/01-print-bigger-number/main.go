package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which receives as input two numbers from the terminal and prints the bigger of the two. If the numbers are equal, print “The numbers are equal”
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	num1Str, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num1, err := strconv.Atoi(strings.TrimSpace(num1Str))
	if err != nil {
		panic(err)
	}
	num2Str, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(strings.TrimSpace(num2Str))
	if err != nil {
		panic(err)
	}
	if num1 > num2 {
		fmt.Printf("\nThe bigger number is %v", num1)
	} else if num2 > num1 {
		fmt.Printf("\nThe bigger number is %v", num2)
	} else {
		fmt.Println("The numbers are equal")
	}
}
