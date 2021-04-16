// https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a

package main

import "fmt"

func Past(h, m, s int) (miliseconds int) {
	return (h*60*60 + m*60 + s) * 1000
}

func main() {
	fmt.Println(Past(0, 1, 1) == 61000)
	fmt.Println(Past(0, 1, 1) == 61000)
	fmt.Println(Past(1, 1, 1) == 3661000)
	fmt.Println(Past(0, 0, 0) == 0)
	fmt.Println(Past(1, 0, 1) == 3601000)
	fmt.Println(Past(1, 0, 0) == 3600000)
}
