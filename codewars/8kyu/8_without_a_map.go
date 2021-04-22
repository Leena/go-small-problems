// https://www.codewars.com/kata/57f781872e3d8ca2a000007e
package main

import "fmt"

func Maps(x []int) (doubled []int) {
	for _, num := range x {
		doubled = append(doubled, num*2)
	}
	return
}

func main() {
	fmt.Println(Maps([]int{1, 2, 3}))          // []int{2, 4, 6}
	fmt.Println(Maps([]int{4, 1, 1, 1, 4}))    // []int{8, 2, 2, 2, 8}
	fmt.Println(Maps([]int{2, 2, 2, 2, 2, 2})) // []int{4, 4, 4, 4, 4, 4}
}
