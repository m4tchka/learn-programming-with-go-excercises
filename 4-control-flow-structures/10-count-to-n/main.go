package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Write a program which takes as input a number N and then generates the numbers from [0; N).
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}
	var res []string
	for i := 0; i < num; i++ {
		res = append(res, fmt.Sprintf("%d", i))
	}
	fmt.Println(strings.Join(res, " "))
}
