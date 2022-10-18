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
Extend the solution to Problem 07* - Guess the number, to interactively ask the user repeatedly to guess the number until they get it right.
In the end, print the number of wrong guesses the user had.
If the number of wrong guesses is less than N/4, print “You were quite lucky!”
If the number of wrong guesses is less than N/2, print “You did alright.”
If the number of wrong guesses is more than or equal to N/2, print “It took you a while…”
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
	var isGuessCorrect bool = false
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(rangeNum)
	var guesses int = 0
	for !isGuessCorrect {
		guessStr, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		guessNum, err := strconv.Atoi(strings.TrimSpace(guessStr))
		if err != nil {
			panic(err)
		}
		if randomNum == guessNum {
			fmt.Printf("You guessed the number correctly!")
			isGuessCorrect = true
		} else {
			fmt.Println("Wrong guess! Try again...")
			guesses++
		}
	}
	var note string
	switch {
	case guesses < rangeNum/4:
		note = "You were quite lucky!"
	case guesses < rangeNum/2:
		note = "You did alright."
	default:
		note = "It took you a while …"
	}
	fmt.Printf("\nYou had %d wrong guesses. %s", guesses, note)
}
