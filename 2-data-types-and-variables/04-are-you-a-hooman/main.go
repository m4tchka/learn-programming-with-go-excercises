package main

import "fmt"

/*
Identify several species around you. Include yourself if you so desire.
For each of them, name them if they aren’t named already & identify if they are human or not.

NOTE: Use boolean variables to complete this task. Don’t just rely on strings

Examples:
Is Gary a hooman?
false
Gary is a cat!

Is Steve a hooman?
true
*/

func main() {
	var Fido bool = false
	fmt.Println("Is Fido a hooman?")
	fmt.Println(Fido)

	var John bool = true
	fmt.Println("Is John a hooman?")
	fmt.Println(John)
}
