package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	Create a program which interactively asks the user to input the name and price of three items on a menu.

Afterwards, print the menu items in the format:
Here’s your restaurant’s menu:
| {item1}		| {price1}$	|
| {item2}		| {price2}$	|
| {item3}		| {price3}$	|
*/
func main() {
	item1 := getInput("Enter the first item's name: ")
	item1Price := parseTerminalFloat(getInput("Enter the first item's price: "))
	item2 := getInput("Enter the second item's name: ")
	item2Price := parseTerminalFloat(getInput("Enter the second item's price: "))
	item3 := getInput("Enter the third item's name: ")
	item3Price := parseTerminalFloat(getInput("Enter the third item's price: "))

	fmt.Println("Here's your restaurant's menu:")
	fmt.Printf("| %-12s | $%-6.2f |\n", item1, item1Price)
	fmt.Printf("| %-12s | $%-6.2f |\n", item2, item2Price)
	fmt.Printf("| %-12s | $%-6.2f |\n", item3, item3Price)

}
func getInput(prompt string) string {
	r := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	terminalInput, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(terminalInput)
}
func parseTerminalFloat(priceStr string) float64 {
	priceFloat, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)
	if err != nil {
		panic(err)
	}
	return priceFloat
}
