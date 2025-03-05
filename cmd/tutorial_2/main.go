package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 1.1
	fmt.Println(floatNum)

	var str string = "Hello, World!"
	fmt.Println(str)

	fmt.Println(utf8.RuneCountInString(str))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var boolVar bool = true
	fmt.Println(boolVar)

	const pi float64 = 3.14159265359
	fmt.Println(pi)
}
