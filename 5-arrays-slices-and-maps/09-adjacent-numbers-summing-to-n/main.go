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
	var result solutions
	for i, n := range numSli {
		if n < target {
			result = checkForward(target, numSli[i:], result)
		} else if n == target {
			var solution subset
			solution.combo = append(solution.combo, n)
			result.validSubsets = append(result.validSubsets, solution)
		} else {
			continue
		}
	}
	fmt.Println("result: ", result)
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
func checkForward(target int, sli []int, res solutions) solutions {
	var sum int
	var possibleSubset subset
	var possibleSolution []int
	for _, num := range sli {
		fmt.Println("Running total: ", sum)
		if sum < target {
			fmt.Println("Number to be added: ", num)
			sum += num
			possibleSolution = append(possibleSolution, num)
		} else if sum == target {
			possibleSubset.combo = append(possibleSubset.combo, possibleSolution...)
			res.validSubsets = append(res.validSubsets, possibleSubset)
			break
		} else {
			fmt.Println("Sum", sum, "has overshot the target", target)
			break
		}
	}
	return res
}

type solutions struct {
	validSubsets []subset
}
type subset struct {
	combo []int
}

/*
[
[1+2+3],
[2+3+1],
[4+1],
[5],
]
*/
