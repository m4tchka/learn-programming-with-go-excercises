package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which takes in two lines of words and excludes all words in the second list from the first one.
*/

func main() {
	sentence := getStrSlice()
	excludedWords := getStrSlice()
	var result []string
	for _, w := range sentence {
		if !checkContains(excludedWords, w) {
			result = append(result, w)
		}
	}
	fmt.Println(strings.Join(result, " "))
}
func getStrSlice() []string {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	strSli := strings.Split(strings.TrimSpace(line), " ")
	return strSli
}
func checkContains(exc []string, wrd string) bool {
	for _, wx := range exc {
		if wrd == wx {
			return true
		}
	}
	return false
}
