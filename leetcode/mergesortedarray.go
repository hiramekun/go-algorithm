package leetcode

import "slices"

// https://leetcode.com/problems/merge-sorted-array?envType=study-plan-v2&envId=top-interview-150
func merge(nums1 []int, m int, nums2 []int, n int) {
	slices.Sort(append(nums1[:m], nums2...))
}
