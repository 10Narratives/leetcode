package leetcode

/*
Given an array nums containing n distinct numbers in the range [0, n],
return the only number in the range that is missing from the array.

Link: https://leetcode.com/problems/missing-number/description/
*/

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func missingNumber(nums []int) int {
	n := len(nums)
	return n*(n+1)/2 - sum(nums)
}
