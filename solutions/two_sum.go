/*https://leetcode.com/problems/two-sum/*/

package solutions

// brute force way O(n^2)
func TwoSum(nums []int, target int) []int {
	for i, e := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (e + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
