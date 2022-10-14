package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Create a program which reads in N numbers from the terminal, one at a line and calculates their median. The numbers are given in sorted order.
The first number shows how many numbers to expect afterwards.
Note: Assume that the elements count is always an odd number. Don’t bother with writing code to calculate the median if you’re given an even count.
Only do that if you want an extra challenge. ;)
*/

func main() {
	length := getInputLength()
	// fmt.Println("Length: ", length)
	median := getInput(length)
	fmt.Println("Median=", median)
}
func getInputLength() int {
	r := bufio.NewReader(os.Stdin)
	lengthStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	length, err := strconv.Atoi(strings.TrimSpace(lengthStr))
	if err != nil {
		panic(err)
	}
	return length
}
func getInput(length int) float64 {
	r := bufio.NewReader(os.Stdin)
	var nums []int
	for i := 0; i < length; i++ {
		numStr, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		inputNum, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			panic(err)
		}
		nums = append(nums, inputNum)
	}
	if length%2 == 0 {
		mid1 := nums[length/2]
		mid2 := nums[(length/2)-1]
		return (float64(mid1+mid2) / 2)
	} else {
		return float64(nums[length/2])
	}
}
