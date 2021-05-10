// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		numberOne := numbers[left]
		numberTwo := numbers[right]
		currentSum := numberOne + numberTwo

		if currentSum == target {
			return []int{left + 1, right + 1}
		}

		if currentSum < target {
			left++
		} else {
			right--
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
