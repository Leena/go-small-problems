// https://www.codewars.com/kata/5667e8f4e3f572a8f2000039

package main

import (
	"fmt"
	s "strings"
)

func Accum(word string) string {
	letters := s.Split(word, "")
	accumulatedLetters := make([]string, len(letters))

	for idx, chr := range letters {
		formatted := (s.ToUpper(chr) + s.Repeat(s.ToLower(chr), idx))
		accumulatedLetters[idx] = formatted
	}

	return s.Join(accumulatedLetters, "-")
}

func main() {
	fmt.Println(Accum("abcd") == "A-Bb-Ccc-Dddd")
	fmt.Println(Accum("RqaEzty") == "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy")
	fmt.Println(Accum("cwAt") == "C-Ww-Aaa-Tttt")
}

