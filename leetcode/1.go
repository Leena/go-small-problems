// https://leetcode.com/problems/two-sum/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numbers := map[int]int{}

	for idx, num := range nums {
		difference := target - num
		value, exists := numbers[difference]
		if exists {
			return []int{value, idx}
		} else {
			numbers[num] = idx
		}
	}
	return nil
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9

	nums2 := []int{3, 2, 4}
	target2 := 6

	nums3 := []int{3, 3}
	target3 := 6

	nums4 := []int{-3, 4, 3, 90}
	target4 := 0

	fmt.Println(twoSum(nums1, target1)) // [0,1]
	fmt.Println(twoSum(nums2, target2)) // [1,2]
	fmt.Println(twoSum(nums3, target3)) // [0,1]
	fmt.Println(twoSum(nums4, target4)) // [0,2]
}
