package main

import "fmt"

/*
Create a program which prints the data for your initial Dark Souls 3 character.

The data should include:
Name - any string
Gender - any string
Age - a number between 0 and 100
Class - any string among {Knight, Mercenary, Warrior, Herald, …, Deprived}
Skin color - a string in format (N, N, N) where N is between 0 and 255 from the RGB spectrum

Use appropriate data types to store all the attributes.
*/
func main() {
	var Name string = "M4tchka"
	var Gender string = "Male"
	var Age int = 22
	var Class string = "Rogue"
	var Skin string = "(0, 255, 255)"
	fmt.Printf("Name: %s\nGender: %s\nAge: %v\nClass: %s\nSkin: %s", Name, Gender, Age, Class, Skin)
}
