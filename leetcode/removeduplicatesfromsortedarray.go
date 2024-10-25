package leetcode

// https://leetcode.com/problems/remove-duplicates-from-sorted-array?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	i := 1
	for k := 1; k < len(nums); k++ {
		if nums[k] != nums[k-1] {
			if i != k {
				nums[i] = nums[k]
			}
			i++
		}
	}
	return i
}
