package leetcode

import "slices"

// https://leetcode.com/problems/remove-element?envType=study-plan-v2&envId=top-interview-150
func removeElement(nums []int, val int) int {
	ret := 0
	for i, e := range nums {
		if e != val {
			ret += 1
		} else {
			nums[i] = 51
		}
	}
	slices.Sort(nums)
	return ret
}
