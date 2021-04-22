// https://www.codewars.com/kata/56bc28ad5bdaeb48760009b0

package main

import "fmt"

func RemoveChar(word string) (formattedString string) {
	for i, char := range word {
		if i == 0 || i == len(word)-1 {
			continue
		} else {
			formattedString += string(char)
		}
	}
	return
}

func main() {
	fmt.Println(RemoveChar("human") == "uma")
	fmt.Println(RemoveChar("country") == "ountr")
	fmt.Println(RemoveChar("place") == "lac")
	fmt.Println(RemoveChar("person") == "erso")
}

/* Alternate Solutions
func RemoveChar(word string) string {
  return word[1:len(word)-1]
}

func RemoveChar(word string) string {
  var newWord = []rune(word)
  return string(newWord[1:len(newWord) - 1])
}
*/
