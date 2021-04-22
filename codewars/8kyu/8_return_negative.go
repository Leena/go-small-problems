// https://www.codewars.com/kata/55685cd7ad70877c23000102

package main

import "fmt"

func MakeNegative(x int) int {
	switch {
	case x < 0:
		return x
	case x > 0:
		return x * -1
	}
	return x
}

func main() {
	fmt.Println(MakeNegative(42) == -42)
	fmt.Println(MakeNegative(0) == 0)
	fmt.Println(MakeNegative(-2) == -2)
}

/* Alternate Solutions
func MakeNegative(x int) int {
	if x < 0 {
		return x
	} else if x > 0 {
		return x * -1
	} else {
		return x
	}
}

func MakeNegative(x int) int {
  if x > 0 {
    x *= -1
  }
  return x
}

func MakeNegative(x int) int {
  if x > 0 {
    return -x
  }
  return x
}
*/
