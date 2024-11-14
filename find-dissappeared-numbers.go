package leetcode

/*
Given an array nums of n integers where nums[i] is in the range [1, n],
return an array of all the integers in the range [1, n] that do not appear in nums.

Link: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
*/

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		pos := nums[i] - 1
		if nums[i] != nums[pos] {
			nums[i], nums[pos] = nums[pos], nums[i]
		} else {
			i++
		}
	}

	disappeared := []int{}
	for i := range len(nums) {
		if nums[i] != i+1 {
			disappeared = append(disappeared, i+1)
		}
	}

	return disappeared
}
