package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Create a program which takes in an input array and finds an adjacent subset of numbers which equals N.
You are given the numbers on the first line and the number N on the second line.
If no numbers sum up to N, print “no solution exists”.
*/

func main() {
	numSli := getIntSli()
	target := getIntInput()
	var result [][]int
	for i, n := range numSli {
		if n < target {
			result = checkForward(target, numSli[i:], result)
		} else if n == target {
			result = append(result, []int{n})
		} else {
			continue
		}
	}
	printResult(result, target)
}
func getIntSli() []int {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	numStrSli := strings.Split((strings.TrimSpace(line)), " ")
	var numSli []int
	for _, numStr := range numStrSli {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}
		numSli = append(numSli, num)
	}
	return numSli
}
func getIntInput() int {
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
func checkForward(target int, sli []int, res [][]int) [][]int {
	var sum int
	var possibleSolution []int
	for _, num := range sli {
		possibleSolution = append(possibleSolution, num)
		sum += num
		if sum < target {
		} else if sum == target {
			res = append(res, possibleSolution)
			break
		} else {
			break
		}
	}
	return res
}
func printResult(res [][]int, t int) {
	fmt.Println()
	if len(res) == 0 {
		fmt.Println("no solution exists")
	} else {
		for _, sli := range res {
			var numStrSli []string
			for _, num := range sli {
				numStr := strconv.Itoa(num)
				numStrSli = append(numStrSli, numStr)
			}
			fmt.Printf("%s = %d\n", strings.Join(numStrSli, "+"), t)
		}
	}
}
