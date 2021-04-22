// https://www.codewars.com/kata/55fd2d567d94ac3bc9000064

package main

import "fmt"

func RowSumOddNumbers(n int) int {
	return n * n * n
}

func main() {
	fmt.Println(RowSumOddNumbers(1) == 1)
	fmt.Println(RowSumOddNumbers(2) == 8)
	fmt.Println(RowSumOddNumbers(13) == 2197)
	fmt.Println(RowSumOddNumbers(19) == 6859)
	fmt.Println(RowSumOddNumbers(41) == 68921)
	fmt.Println(RowSumOddNumbers(42) == 74088)
	fmt.Println(RowSumOddNumbers(74) == 405224)
	fmt.Println(RowSumOddNumbers(86) == 636056)
	fmt.Println(RowSumOddNumbers(93) == 804357)
	fmt.Println(RowSumOddNumbers(101) == 1030301)
}
