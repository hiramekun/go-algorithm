package leetcode

// https://leetcode.com/problems/remove-duplicates-from-sorted-array?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}
	return len(nums[:left+1])
}
