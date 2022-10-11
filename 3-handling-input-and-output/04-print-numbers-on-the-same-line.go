package main

/*
Create a program which initializes three integer variables with the numbers 2, 3, 4.
Print them on the same line using a suitable routine.
*/
import "fmt"

func main() {
	n1, n2, n3 := 2, 3, 4
	fmt.Println(n1, n2, n3)
}

/*
go build 04-print-numbers-on-the-same-line.go
./04-print-numbers-on-the-same-line.exe
mv 04-print-numbers-on-the-same-line.exe 04.exe
./04.exe
rm 04.exe
*/
