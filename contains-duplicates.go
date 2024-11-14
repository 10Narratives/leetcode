package leetcode

/*
Given an integer array nums, return true if any value appears at least
twice in the array, and return false if every element is distinct.

Link: https://leetcode.com/problems/contains-duplicate/description/
*/

func containsDuplicates(nums []int) bool {
	check_result := false
	metNums := make(map[int]bool, 0)

	for _, num := range nums {
		if _, contains := metNums[num]; contains {
			check_result = contains
			break
		} else {
			metNums[num] = true
		}
	}

	return check_result
}
