// https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097

package main

import "fmt"

func century(year int) (century int) {
	remainder := year % 100

	switch remainder {
	case 0:
		century = year / 100
	default:
		century = (year / 100) + 1
	}
	return
}

func main() {
	fmt.Println(century(1705) == 18)
	fmt.Println(century(1900) == 19)
	fmt.Println(century(1601) == 17)
}

/* Alternate Solutions
func century(year int) int {
  return (year + 99)/100
}

func century(year int) int {
 if year % 100 != 0  { year += 100 }
 return year / 100
}

*/
