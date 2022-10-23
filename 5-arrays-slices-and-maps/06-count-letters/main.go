package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

/*
Write a program which takes in some text and counts the number of letters in it.
The counting should be case-insensitive.
*/

func main() {
	sentence := getSentence()
	letters := map[rune]int{}
	for _, char := range sentence {
		if !unicode.IsLetter(char) {
			continue
			//Necessary because without unicode.Isletter, spaces and special characters would also be added to the map
		}
		letters[char]++
		//Alternative if sentence was not all lowercase: letters[unicode.toLower(char)]++
	}
	keys := make([]int, 0, len(letters)) // Make a slice to contain all the keys of the map, based on the length of the map
	for k := range letters {             // For each unique rune key in the letters map, convert it to an integer and append it to the slice. (value would be second returned value but unneccessary here)
		keys = append(keys, int(k))
	}
	sort.Ints(keys)          // Sort the slice of runes that were converted to integers numerically ascending (which is also alphabetical order for runes)
	for _, k := range keys { // Loop through the slice of ints (which correspond to a rune)
		fmt.Printf("%c: %d\n", k, letters[rune(k)]) // Print each int and the corresponding value of that key (after converting back to a rune) in the letters map.
	}
}
func getSentence() string {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	sntc := strings.ToLower(strings.TrimSpace(line))
	//NOTE: Alternatively, lowercase conversion can be done in the for loop using unicode.ToLower(rune)
	return sntc
}

/*

if _,ok := letters[letter];!ok {
	letters[letter] = 1
} else {
	letters[letter] ++
}
*/

/*
	for key, val := range letters {
		fmt.Printf("%c: %d\n", key, val)
	}
*/
