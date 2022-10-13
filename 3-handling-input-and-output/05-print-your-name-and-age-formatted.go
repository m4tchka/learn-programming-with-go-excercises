package main

/*
Create a program, which initializes two variables - name & age.

Print them on the terminal in the following format:
Hi, my name is {name}!
I'm {age} years old and how old are you?
*/
import "fmt"

func main() {
	var name string = "Matchka"
	age := 22
	fmt.Printf("Hi,my name is %s!\nI'm %d years old and how old are you ?", name, age)
}
