/*
 * @lc app=leetcode.cn id=523 lang=golang
 * @lcpr version=30403
 *
 * [523] 连续的子数组和
 */

package leetcode_solutions

import "testing"

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	// 思路 preSum
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	val2idx := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		val := preSum[i] % k
		if idx, ok := val2idx[val]; ok {
			if i-idx >= 2 {
				return true
			}
		} else {
			val2idx[val] = i
		}
	}

	return false
}

// @lc code=end

func TestContinuousSubarraySum(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [23,2,4,6,7]\n6\n
// @lcpr case=end

// @lcpr case=start
// [23,2,6,4,7]\n6\n
// @lcpr case=end

// @lcpr case=start
// [23,2,6,4,7]\n13\n
// @lcpr case=end

*/
