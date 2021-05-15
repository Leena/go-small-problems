// https://leetcode.com/problems/roman-to-integer/description/

package main

import "fmt"

func romanToInt(s string) (sum int) {
	romanToDigits := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && romanToDigits[s[i]] < romanToDigits[s[i+1]] {
			sum -= romanToDigits[s[i]]
		} else {
			sum += romanToDigits[s[i]]
		}
	}

	return
}

func main() {
	num1 := "III"
	target1 := 3

	num2 := "IV"
	target2 := 4

	num3 := "IX"
	target3 := 9

	num4 := "MCMXCIV"
	target4 := 1994

	fmt.Println(romanToInt(num1) == target1)
	fmt.Println(romanToInt(num2) == target2)
	fmt.Println(romanToInt(num3) == target3)
	fmt.Println(romanToInt(num4) == target4)
}
