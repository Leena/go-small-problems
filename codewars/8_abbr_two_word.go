// https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3

package main

import (
	"fmt"
	"strings"
)

func AbbrevName(fullName string) (abrev string) {
	listNames := strings.Split(strings.ToUpper(fullName), " ")
	initials := make([]string, len(listNames))

	for index, name := range listNames {
		initials[index] = string(name[0])
	}

	return strings.Join(initials, ".")
}

func main() {
	fmt.Println(AbbrevName("Sam Harris") == "S.H")
	fmt.Println(AbbrevName("Patrick Feenan") == "P.F")
	fmt.Println(AbbrevName("Evan Cole") == "E.C")
	fmt.Println(AbbrevName("P Favuzzi") == "P.F")
	fmt.Println(AbbrevName("David Mendieta") == "D.M")
}

/* Alternate Solution
func AbbrevName(name string) string{
  words := strings.Split(name, " ")
  return strings.ToUpper(string(words[0][0])) + "." + strings.ToUpper(string(words[1][0]))
}

func AbbrevName(name string) string{
  var parts []string
  for _, part := range strings.Split(name, " ") {
    parts = append(parts, strings.ToUpper(part[:1]))
  }
  return strings.Join(parts, ".")
}
*/
