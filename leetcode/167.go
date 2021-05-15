// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {
		n1 := numbers[l]
		n2 := numbers[r]
		cs := n1 + n2

		if cs == target {
			return []int{l + 1, r + 1}
		}

		if cs < target {
			l++
		} else {
			r--
		}
	}
	return nil
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9

	nums2 := []int{2, 3, 4}
	target2 := 6

	nums3 := []int{-1, 0}
	target3 := -1

	fmt.Println(twoSum(nums1, target1)) // [1,2]
	fmt.Println(twoSum(nums2, target2)) // [1,3]
	fmt.Println(twoSum(nums3, target3)) // [1,2]
}
