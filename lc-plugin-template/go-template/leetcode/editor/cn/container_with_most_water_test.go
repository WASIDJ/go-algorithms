/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=30403
 *
 * [11] 盛最多水的容器
 */

package leetcode_solutions

import "testing"

// @lc code=start
func maxArea(height []int) int {
	// 思路:双指针: 使用两个指针分别指向数组的两端，计算当前指针所指元素之间的面积，并更新最大面积。然后根据指针所指元素的高度，移动较短的指针向内移动，以期望找到更大的面积。
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		currentArea := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, currentArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestContainerWithMostWater(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

*/
