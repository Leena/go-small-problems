// https://www.codewars.com/kata/5656b6906de340bd1b0000ac

package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) (uniqueChars string) {
	merged := s1 + s2
	stringLetters := strings.Split(merged, "")
	sort.Strings(stringLetters)

	for _, letter := range stringLetters {
		if !strings.Contains(uniqueChars, letter) {
			uniqueChars += letter
		}
	}
	return
}

func main() {
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere") == "aehrsty")
	fmt.Println(TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding") == "abcdefghilnoprstu")
}
