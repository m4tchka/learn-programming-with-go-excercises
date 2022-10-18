package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which determines which game a character belongs to from an input character name.

The characters which will be given by game are:
For The Witcher 3:
Geralt
Ciri
Yennefer

For StarCraft 2:
Aldaris
Jim Raynor
Kerrigan
Zeratul

For WarCraft 3:
Arthas
Illidan
Sylvanas

Based on the input name, print {character name} is part of the game {game name}!
If you get a character which is not in this list, print I don’t recognize this character…
To complete this exercise, use a switch-case statement with fallthroughs.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	c, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	char := strings.TrimSpace(c)
	switch char {
	case "Geralt", "Ciri", "Yennefer":
		fmt.Printf("\n%s is part of the game Witcher 3!", char)
	case "Aldaris", "Jim Raynor", "Kerrigan", "Zeratul":
		fmt.Printf("\n%s is part of the game StarCraft 2 !", char)
	case "Arthas", "Illidan", "Sylvanas":
		fmt.Printf("\n%s is part of the game WarCraft 3!", char)
	default:
		fmt.Printf("\nI don't recognize this character…")
	}
}
