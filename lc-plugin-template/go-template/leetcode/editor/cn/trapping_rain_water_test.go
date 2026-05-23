/*
 * @lc app=leetcode.cn id=42 lang=golang
 * @lcpr version=30403
 *
 * [42] 接雨水
 */

package leetcode_solutions

import "testing"

// @lc code=start
func trap(height []int) int {
	// 思路1:暴力解法 n^2
	// 思路2:备忘录 n
	// 思路3:双指针 n
	// Q:此时的 l_max 是 left 指针左边的最高柱子，但是 r_max 并不一定是 left 指针右边最高的柱子，这真的可以得到正确答案吗？
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	res := 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}

	return res
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

func TestTrappingRainWater(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "case2",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trap(tt.height)
			if result != tt.expected {
				t.Errorf("trap(%v) = %d, want %d", tt.height, result, tt.expected)
			}
		})
	}
}

/*
// @lcpr case=start
// [0,1,0,2,1,0,1,3,2,1,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,0,3,2,5]\n
// @lcpr case=end

*/
