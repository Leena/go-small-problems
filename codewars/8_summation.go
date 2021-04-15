// https://www.codewars.com/kata/55d24f55d7dd296eb9000030

package main

import "fmt"

func Summation(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func main() {
	fmt.Println(Summation(1) == 1)
	fmt.Println(Summation(8) == 36)
}
