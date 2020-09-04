package leetcode

// brute force way O(n^2)
func twoSum(nums []int, target int) []int {
	for i, e := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (e + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
