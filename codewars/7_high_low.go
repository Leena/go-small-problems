// https://www.codewars.com/kata/554b4ac871d6813a03000035

package main

import (
	"fmt"
	"sort"
	stcv "strconv"
	s "strings"
)

func highAndLow(nums string) string {
	sliceStrNums := s.Split(nums, " ")
	numsArray := make([]int, len(sliceStrNums))
	for idx, strNum := range sliceStrNums {
		num, _ := stcv.Atoi(strNum)
		numsArray[idx] = num
	}

	sort.Ints(numsArray)
	high := numsArray[len(numsArray)-1]
	low := numsArray[0]

	return stcv.Itoa(high) + " " + stcv.Itoa(low)
}

func main() {
	fmt.Println(highAndLow("1 2 3 4 5") == "5 1")
	fmt.Println(highAndLow("1 2 -3 4 5") == "5 -3")
	fmt.Println(highAndLow("1 9 3 4 -5") == "9 -5")
	fmt.Println(highAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4") == "42 -9")
}
