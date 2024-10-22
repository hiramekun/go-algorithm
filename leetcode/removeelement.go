package leetcode

import "slices"

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
