package main

import (
	"fmt"
	"strings"
)

/*
Write a program which buys the best possible choice, among three games you love, based on an input number, which represents the cash your parents have given you.
There are three games you love, in this order:
StarCraft 2 - 30$
Cyberpunk 2077 - 130$
The Witcher 3 - 60$

Given an input budget N, decide which games you’re going to buy. If you can afford all of them, then have them all!

If you can afford only 2 out of three, then make your preference in this order based on the budget you have.

The output should be in the following format:
Here’s what I bought:
{game1}, {game2}, {game 3}

If you bought less than three games, print the games you did buy in this order.
If you couldn’t buy anything, print I couldn’t buy anything!
*/

func main() {
	budget := 90
	var games []string
	if budget >= 30 {
		budget -= 30
		// sc2 = true
		games = append(games, "StarCraft2")
	} else {
		fmt.Println("I couldn't buy anything!")
	}
	if budget >= 130 {
		budget -= 130
		// cp2077 = true
		games = append(games, "Cyberpunk2077")
	}
	if budget >= 60 {
		budget -= 60
		// tw3 = true
		games = append(games, "The Witcher 3")
	}
	if len(games) == 0 {
		fmt.Println("I couldn’t buy anything!")
	} else {
		fmt.Println("Here's what I bought:")
		fmt.Println(strings.Join(games, ", "))
	}
}
