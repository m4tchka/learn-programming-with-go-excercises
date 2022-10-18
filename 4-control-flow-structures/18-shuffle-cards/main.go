package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
You are given a set of standard playing cards. They are represented by the numbers 2 through 10 and the characters A, J, Q, K.
Write a program to shuffle the cards in a random order.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	deck := strings.Split(strings.TrimSpace(line), " ")
	fmt.Println("Unshuffled deck: ", deck)
	var newDeck []string
	for len(deck) > 0 {
		cutIndex := rand.Intn(len(deck))
		cardMovedToEnd := deck[cutIndex]
		deck = removeCard(deck, cutIndex)
		newDeck = append(newDeck, cardMovedToEnd)
	}
	fmt.Printf("Shuffled deck: %s", newDeck)
}

func removeCard(deck []string, cut int) []string {
	return append(deck[:cut], deck[cut+1:]...)
}
