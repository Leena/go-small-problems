// https://www.codewars.com/kata/563b662a59afc2b5120000c6

package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) (years int) {
	for p0 < p {
		p0 = p0 + int(float64(p0)*(percent/100)) + aug
		years++
	}
	return
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000) == 15)
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000) == 10)
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000) == 94)
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000) == 151)
}
