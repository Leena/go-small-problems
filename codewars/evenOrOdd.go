// Source: https://www.codewars.com/kata/53da3dbb4a5168369a0000fe

package main

import "fmt"

func EvenOrOdd(number int) string {

	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	test1 := EvenOrOdd(-1)
	test2 := EvenOrOdd(-2)
	test3 := EvenOrOdd(1)
	test4 := EvenOrOdd(2)
	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	fmt.Println(test4)
}
