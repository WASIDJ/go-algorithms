/*
 * @lc app=leetcode.cn id=1124 lang=golang
 * @lcpr version=30403
 *
 * [1124] 表现良好的最长时间段
 */

package leetcode_solutions

import "testing"

// @lc code=start
func longestWPI(hours []int) int {
	// 思路 感觉是 525的变体 变成 nums[i]>8 sum++ nums[i]<8 sum--
	n := len(hours)
	preSum := make([]int, n+1)
	valToIndex := make(map[int]int)
	res := 0
	for i := 1; i <= n; i++ {
		if hours[i-1] > 8 {
			preSum[i] = preSum[i-1] + 1
		} else {
			preSum[i] = preSum[i-1] - 1
		}
		if _, found := valToIndex[preSum[i]]; !found {

			valToIndex[preSum[i]] = i

		}
		if preSum[i] > 0 {

			// preSum[i] 为正，说明 hours[0..i-1] 都是「表现良好的时间段」

			res = max(res, i)

		} else {

			// preSum[i] 为负，需要寻找一个 j 使得 preSum[i] - preSum[j] > 0

			// 考虑到我们的 preSum 数组每两个相邻元素的差的绝对值都是 1 且 j 应该尽可能小，

			// 那么只要找到 preSum[j] == preSum[i] - 1，nums[j+1..i] 就是一个「表现良好的时间段」

			if j, found := valToIndex[preSum[i]-1]; found {

				res = max(res, i-j)

			}
		}
	}
	return res
}
func max3(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

func TestLongestWellPerformingInterval(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [9,9,6,0,6,6,9]\n
// @lcpr case=end

// @lcpr case=start
// [6,6,6]\n
// @lcpr case=end

*/
