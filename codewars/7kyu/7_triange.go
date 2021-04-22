// https://www.codewars.com/kata/56606694ec01347ce800001b

package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	return (((a + b) > c) && ((a + c) > b) && ((c + b) > a))
}

func main() {
	fmt.Println(IsTriangle(5, 1, 5) == true)
	fmt.Println(IsTriangle(5, 1, 2) == false)
}
