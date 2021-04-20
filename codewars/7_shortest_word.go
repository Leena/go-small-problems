// https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9

package main

import (
	"fmt"
	"strings"
)

func findShort(s string) int {
	arrayWords := strings.Split(s, " ")
	short := len(arrayWords[0])

	for _, word := range arrayWords {
		if len(word) < short {
			short = len(word)
		}
	}
	return short
}

func main() {
	fmt.Println(findShort("should test that the solution returns the correct value") == 3)
	fmt.Println(findShort("bitcoin take over the world maybe who knows perhaps") == 3)
	fmt.Println(findShort("turns out random test cases are easier than writing out basic ones") == 3)
}

/* Alternative Solution
func FindShort(s string) int {
  shortest := len(s)
  for _, word := range strings.Split(s, " ") {
    if len(word) < shortest {
      shortest = len(word)
    }
  }
  return shortest
}
*/
