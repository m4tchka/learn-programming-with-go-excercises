package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which reads commands from the terminal & performs operations on a set of numbers.
The commands are:
add <num> - add a new number to the collection of numbers
inc <index> - increment a number at a given index (zero-based)
sub <index> <num> - subtract num from the number at the given index
mul <index> <num> - multiply the number at the given index by num
acc <index> <num> - accumulate num to the number at the given index
show - print the current value of all numbers at a single line
exit - stop program execution
If the provided index is outside the range of numbers you currently have, print “invalid index”. If the operation was executed successfully, print “OK”.
*/

func main() {
	var numSlice []int
	getInputOption(numSlice)
}

func getInputOption(nS []int) {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSpace(line)
	inputStrSlice := strings.Split(line, " ")
	switch inputStrSlice[0] {
	case "add":
		nS = add(nS, inputStrSlice[1])
		getInputOption(nS)
	case "inc":
		if strToInt(inputStrSlice[1]) >= len(nS) {
			fmt.Println("Invalid index")
			getInputOption(nS)
		} else {
			nS = inc(nS, inputStrSlice[1])
			getInputOption(nS)
		}
	case "sub":
		if strToInt(inputStrSlice[1]) >= len(nS) {
			fmt.Println("Invalid index")
			getInputOption(nS)
		} else {
			nS = sub(nS, inputStrSlice[1], inputStrSlice[2])
			getInputOption(nS)
		}
	case "mul":
		if strToInt(inputStrSlice[1]) >= len(nS) {
			fmt.Println("Invalid index")
			getInputOption(nS)
		} else {
			nS = mul(nS, inputStrSlice[1], inputStrSlice[2])
			getInputOption(nS)
		}
	case "acc":
		if strToInt(inputStrSlice[1]) >= len(nS) {
			fmt.Println("Invalid index")
			getInputOption(nS)
		} else {
			nS = acc(nS, inputStrSlice[1], inputStrSlice[2])
			getInputOption(nS)
		}
	case "show":
		show(nS)
		getInputOption(nS)
	case "exit":
		break
	}
}
func strToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
func add(nS []int, newNumStr string) []int {
	// newNum, err := strconv.Atoi(newNumStr)
	// if err != nil {
	// 	panic(err)
	// }
	nS = append(nS, strToInt(newNumStr))
	fmt.Println("OK")
	return nS
}
func inc(nS []int, indexStr string) []int {
	// index, err := strconv.Atoi(indexStr)
	// if err != nil {
	// 	panic(err)
	// }
	nS[strToInt(indexStr)]++
	fmt.Println("OK")
	return nS
}
func sub(nS []int, indexStr string, numStr string) []int {
	// index, err := strconv.Atoi(indexStr)
	// if err != nil {
	// 	panic(err)
	// }
	// num, err := strconv.Atoi(numStr)
	// if err != nil {
	// 	panic(err)
	// }
	nS[strToInt(indexStr)] -= strToInt(numStr)
	fmt.Println("OK")
	return nS
}
func mul(nS []int, indexStr string, numStr string) []int {
	// index, err := strconv.Atoi(indexStr)
	// if err != nil {
	// 	panic(err)
	// }
	// num, err := strconv.Atoi(numStr)
	// if err != nil {
	// 	panic(err)
	// }
	nS[strToInt(indexStr)] *= strToInt(numStr)
	fmt.Println("OK")
	return nS
}
func acc(nS []int, indexStr string, numStr string) []int {
	// index, err := strconv.Atoi(indexStr)
	// if err != nil {
	// 	panic(err)
	// }
	// num, err := strconv.Atoi(numStr)
	// if err != nil {
	// 	panic(err)
	// }
	nS[strToInt(indexStr)] += strToInt(numStr)
	fmt.Println("OK")
	return nS
}
func show(nS []int) {
	var nSStr []string
	for _, num := range nS {
		nSStr = append(nSStr, fmt.Sprintf("%v", num))
	}
	fmt.Println(strings.Join(nSStr, " "))
}
