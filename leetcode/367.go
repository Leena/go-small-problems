// https://leetcode.com/problems/valid-perfect-square/

package main

import "fmt"

func isPerfectSquare(num int) bool {
	l := 1
	r := num

	for l < r {
		m := l + (r-l)/2
		sm := m * m

		switch {
		case sm == num:
			return true
		case sm < num:
			l = m + 1
		case sm > num:
			r = m - 1
		}
	}

	if l*l == num {
		return true
	}

	return false
}

func main() {
	num1 := 16
	num2 := 14

	fmt.Println(isPerfectSquare(num1) == true)
	fmt.Println(isPerfectSquare(num2) == false)
}
