package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Learn about Selection Sort & sort an input array of numbers using the algorithm.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	numStrSli := strings.Split(strings.TrimSpace(line), " ")
	if err != nil {
		panic(err)
	}
	var unsortedNumSli []int
	for i := range numStrSli {
		num, err := strconv.Atoi(numStrSli[i])
		if err != nil {
			panic(err)
		}
		unsortedNumSli = append(unsortedNumSli, num)
	}

	sortedNumSli := selectionSort(unsortedNumSli)
	fmt.Printf("Sorted: %v", sortedNumSli)

	var sortedNumStrSli []string
	for _, num := range sortedNumSli {
		sortedNumStrSli = append(sortedNumStrSli, strconv.Itoa(num))
	}
	fmt.Println(strings.Join(sortedNumStrSli, " "))
}
func selectionSort(unsortedNumSli []int) []int {
	var sortedNumSli []int
	for _, num := range unsortedNumSli {
		fmt.Println("Num: ", num)
		lowest := unsortedNumSli[0]
		// lowestIndex := 0
		if num < lowest {
			lowest = num
			// lowestIndex = i
		}
		fmt.Println("Lowest: ", lowest)
		sortedNumSli = append(sortedNumSli, lowest)
	}
	return sortedNumSli
}

// func remove(intSli []int, index int) []int {
// 	intSli = append(intSli[:index], intSli[index+1:]...)
// 	return intSli
// }
