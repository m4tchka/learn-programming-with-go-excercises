package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which takes an input number in the range [0; 1000) and prints it as a word.
*/

func main() {
	r := bufio.NewReader(os.Stdin)
	numStr, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	strSlc := strings.Split(strings.TrimSpace(numStr), "")
	// fmt.Println(strSlc)
	var word string

	if len(strSlc) == 3 {
		word += fmt.Sprint(checkHundreds(strSlc[0]))
		word += fmt.Sprint(checkTens(strSlc[1], strSlc[2]))
	} else if len(strSlc) == 2 {
		word += fmt.Sprint(checkTens(strSlc[0], strSlc[1]))
	} else {
		word += fmt.Sprint(checkOnes(strSlc[0]))
	}

	fmt.Println(strings.TrimSpace(word))
}

func checkHundreds(hundred string) string {
	switch hundred {
	case "1":
		return "one hundred and "
	case "2":
		return "two hundred and "
	case "3":
		return "three hundred and "
	case "4":
		return "four hundred and "
	case "5":
		return "five hundred and "
	case "6":
		return "six hundred and "
	case "7":
		return "seven hundred and "
	case "8":
		return "eight hundred and "
	case "9":
		return "nine hundred and "
	default:
		return ""
	}
}
func checkTens(ten, one string) string {
	switch ten {
	case "1":
		switch one {
		case "1":
			return "eleven"
		case "2":
			return "twelve"
		case "3":
			return "thirteen"
		case "4":
			return "fourteen"
		case "5":
			return "fifteen"
		case "6":
			return "sixteen"
		case "7":
			return "seventeen"
		case "8":
			return "eighteen"
		case "9":
			return "nineteen"
		default:
			return "ten"
		}
	case "2":
		return fmt.Sprintf("twenty %s", checkOnes(one))
	case "3":
		return fmt.Sprintf("thirty %s", checkOnes(one))
	case "4":
		return fmt.Sprintf("forty %s", checkOnes(one))
	case "5":
		return fmt.Sprintf("fifty %s", checkOnes(one))
	case "6":
		return fmt.Sprintf("sixty %s", checkOnes(one))
	case "7":
		return fmt.Sprintf("seventy %s", checkOnes(one))
	case "8":
		return fmt.Sprintf("eighty %s", checkOnes(one))
	case "9":
		return fmt.Sprintf("ninety %s", checkOnes(one))
	default:
		return checkOnes(one)
	}
}
func checkOnes(one string) string {
	switch one {
	case "1":
		return "one"
	case "2":
		return "two"
	case "3":
		return "three"
	case "4":
		return "four"
	case "5":
		return "five"
	case "6":
		return "six"
	case "7":
		return "seven"
	case "8":
		return "eight"
	case "9":
		return "nine"
	default:
		return ""
	}
}
