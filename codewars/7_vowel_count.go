// https://www.codewars.com/kata/54ff3102c1bad923760001f3

package main

import (
	"fmt"
	"strings"
)

func GetCount(word string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	letters := strings.Split(word, "")
	for _, char := range letters {
		if contains(vowels, string(char)) {
			count++
		}
	}
	return
}

func contains(slice []string, str string) bool {
	for _, element := range slice {
		if element == str {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(GetCount("abracadabra") == 5)
}

/* Alternate Solution
func GetCount(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}

func GetCount(strn string) int {
  count := 0
  vowels := []string{"a", "e", "i", "o", "u"}

  for _, vowel := range vowels {
    count += strings.Count(strn, vowel)
  }
  return count
}

func GetCount(str string) (count int) {
  r := regexp.MustCompile("[aeiou]")
  vowels := r.FindAllString(str, -1)
  return len(vowels)
}
*/
