package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which calculates the roots of a quadratic equation of the form ax^2 + bx + c = 0. More info.
On the terminal, you are given three numbers - a, b and c which form the variables in the equation. If no real roots exist, print no real roots exist!
*/

func main() {
	a := getInput()
	b := getInput()
	c := getInput()
	ds := getDiscriminant(a, b, c)
	if ds < 0 {
		fmt.Println("No real roots exist")
	} else if ds == 0 {
		x := getPosSolution(a, b, c)
		fmt.Printf("x = %.2f", x)
	} else {
		x1 := getPosSolution(a, b, c)
		x2 := getNegSolution(a, b, c)
		fmt.Printf("x1 = %.2f\nx2 = %.2f", x1, x2)
	}
}
func getInput() float64 {
	r := bufio.NewReader(os.Stdin)
	numStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	numFlt, err := strconv.ParseFloat(strings.TrimSpace(numStr), 64)
	if err != nil {
		panic(err)
	}
	// fmt.Println("numFlt: ", numFlt)
	// fmt.Printf("numFlt is type %T", numFlt)
	return numFlt
}
func getDiscriminant(a, b, c float64) float64 {
	ds := (b * b) - (4 * a * c)
	fmt.Printf("discrimant: %v\n", ds)
	return ds
}
func getPosSolution(a, b, c float64) float64 {
	x := (-b + math.Sqrt((b*b)-(4*a*c))) / (2 * a)
	return x
}
func getNegSolution(a, b, c float64) float64 {
	x := (-b - math.Sqrt((b*b)-(4*a*c))) / (2 * a)
	return x
}
