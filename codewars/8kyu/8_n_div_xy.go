// https://www.codewars.com/kata/5545f109004975ea66000086

package main

import "fmt"

func IsDivisible(n, x, y int) bool {
	if n%y == 0 && n%x == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsDivisible(3, 3, 1) == true)
	fmt.Println(IsDivisible(12, 3, 4) == true)
	fmt.Println(IsDivisible(8, 3, 4) == false)
	fmt.Println(IsDivisible(48, 3, 4) == true)
	fmt.Println(IsDivisible(100, 5, 3) == false)
}

/* Alternate Solution

func IsDivisible(n, x, y int) bool {
    return n % x == 0 && n % y == 0
}

func IsDivisible(n, x, y int) bool {
    return n % x + n % y == 0
}
*/
