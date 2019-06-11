/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
func twoSum(nums []int, target int) []int {
	var result []int
    for i := 0; i < len(nums) -1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
	}

	return result
}

