package leetcode

/*
Given a non-empty array of integers nums, every element appears twice except for one.
Find that single one.

You must implement a solution with a linear runtime complexity and use only constant
extra space.

Link: https://leetcode.com/problems/single-number/description/
*/

func singleNumber(nums []int) int {
	mask := 0

	for _, num := range nums {
		mask ^= num
	}

	return mask
}
