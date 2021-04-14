//https://www.codewars.com/kata/515e271a311df0350d00000f

package main

import (
	"fmt"
	"math"
)

func SquareSum(numbers []int) (square int) {
	for _, v := range numbers {
		square += int(math.Pow(float64(v), 2))
	}
	return
}

func main() {
	test1 := []int{1, 2}
	test2 := []int{0, 3, 4, 5}
	fmt.Println(SquareSum(test1) == 5)
	fmt.Println(SquareSum(test2) == 50)
}
