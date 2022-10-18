package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
Write a program which generates a random number in a given range and the user has to guess what number your program has in mind (or in memory, to be more specific…)
It should read two input numbers from the terminal:
N - the range of numbers to use
X - the number the user has guessed

Then, the program should generate a random number in the range [0, N) and check if the number the user guessed is equal to the number the program generated.
If so, print:
You guessed the number correctly!
If not, print:
That’s not the number I had in mind. You guessed {X} but my number was {N}
*/
func main() {
	r := bufio.NewReader(os.Stdin)
	rangeStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	rangeNum, err := strconv.Atoi(strings.TrimSpace(rangeStr))
	if err != nil {
		panic(err)
	}
	guessStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	guessNum, err := strconv.Atoi(strings.TrimSpace(guessStr))
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(rangeNum)
	if randomNum == guessNum {
		fmt.Printf("\nYou guessed the number correctly!")
	} else {
		fmt.Printf("\nThat’s not the number I had in mind. You guessed %d but my number was %d", guessNum, randomNum)
	}
}
