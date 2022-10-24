package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which finds the longest word in a given text.
If multiple longest words exist, print the leftmost one.
*/

func main() {
	sentenceSli := getSentence()
	var longest string
	for _, v := range sentenceSli {
		if len(v) > len(longest) {
			longest = v
		}
	}
	fmt.Println(longest)
}
func getSentence() []string {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	sntSli := strings.Split(strings.TrimSpace(line), " ")
	return sntSli
}
