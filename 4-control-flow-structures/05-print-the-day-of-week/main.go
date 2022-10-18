package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which prints the day of the week given an input integer, which represents the day you have to print.
1 is for Monday, 7 is for Sunday.
If you’re given a number which is not in the bounds [1; 7], then print I don’t recognize that day of the week!
Hint: use a switch-case statement!
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	numStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(numStr))
	if err != nil {
		panic(err)
	}
	switch num {
	case 1:
		fmt.Println("Monday")
		break
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("I don’t recognize that day of the week!")
	}
}
