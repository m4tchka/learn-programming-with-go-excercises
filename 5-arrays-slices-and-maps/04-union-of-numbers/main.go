package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which takes in two lines of numbers and outputs the union of the two lists. There should be no duplicates.
Note: The order of the numbers in the final output doesn’t matter
*/

func main() {
	sli1 := getNumSlice()
	sli2 := getNumSlice()
	var sliU []string
	sliU = append(sliU, sli1...)
	for _, elem := range sli2 {
		if !checkContains(sliU, elem) {
			sliU = append(sliU, elem)
		} else {
			continue
		}
	}
	fmt.Println(strings.Join(sliU, " "))
}

func getNumSlice() []string {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	numStrSli := strings.Split(strings.TrimSpace(line), " ")
	return numStrSli
}
func checkContains(strSli []string, e string) bool {
	for _, elem := range strSli {
		if elem == e {
			return true
		}
	}
	return false
}
