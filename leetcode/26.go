// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package main

import "fmt"

func removeDuplicates(nums []int) (w int) {
	for r, n := range nums {
		if r == 0 || n != nums[r-1] {
			nums[w] = nums[r]
			w++
		}
	}
	return
}

func main() {
	input1 := []int{1, 1, 2}
	input2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	fmt.Println(removeDuplicates(input1) == 2) // 2 unique elements [1, 2]
	fmt.Println(removeDuplicates(input2) == 5) // 5 unique elements [0, 1, 2, 3, 4]
}
