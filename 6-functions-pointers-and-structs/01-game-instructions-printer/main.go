package main

import "fmt"

/*
Create a program which has a function printGameInstructions, which doesn’t take any parameters and doesn’t return a result. It should print the instructions for the hangman game.
*/
func main() {
	printGameInstructions()
}
func printGameInstructions() {
	fmt.Println(`
1. Have the player select a letter of the alphabet.
2. If the letter is contained in the word/phrase, 
   all the the letters equal to it are revealed.
3. If the letter is not contained in the word/phrase, a portion of the hangman is added.
4. The game continues until:
	a. the word/phrase is guessed (all letters are revealed) – WINNER or,
	b. all the parts of the hangman are displayed – LOSER
	`)
}
