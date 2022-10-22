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
	var sortedNumStrSli []string
	for _, num := range sortedNumSli {
		sortedNumStrSli = append(sortedNumStrSli, strconv.Itoa(num))
	}
	fmt.Println(strings.Join(sortedNumStrSli, " "))
}
func selectionSort(unsortedNumSli []int) []int {
	// Loop through unsortedNumSli
	// Record the min and its index
	// Append min to new array
	// Remove min from old array
	for startIndex := 0; startIndex < len(unsortedNumSli); startIndex++ {
		minIndex := startIndex
		for i := startIndex; i < len(unsortedNumSli); i++ {
			if unsortedNumSli[i] < unsortedNumSli[minIndex] {
				minIndex = i
			}
		}
		fmt.Println("On loop", startIndex, "the lowest number is:", unsortedNumSli[minIndex], " at index ", minIndex)
		unsortedNumSli[startIndex], unsortedNumSli[minIndex] = unsortedNumSli[minIndex], unsortedNumSli[startIndex]
	}
	return unsortedNumSli
}
