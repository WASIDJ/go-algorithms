/*
 * @lc app=leetcode.cn id=560 lang=golang
 * @lcpr version=30403
 *
 * [560] 和为 K 的子数组
 */

package leetcode_solutions

import "testing"

// @lc code=start
func subarraySum(nums []int, k int) int {
	// preSum
	n := len(nums)
	preSum := make(map[int]int, n+1)
	preSum[0] = 0
	count := make(map[int]int)
	count[0] = 1
	res := 0
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
		need := preSum[i] - k
		if _, ok := count[need]; ok {
			res = res + count[need]
		}
		count[preSum[i]]++
	}
	return res
}

// @lc code=end

func TestSubarraySumEqualsK(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{nums: []int{1, 2, 3}, k: 3, want: 2},
		{nums: []int{1, 1, 1}, k: 2, want: 2},
	}

	for _, tt := range tests {
		if got := subarraySum(tt.nums, tt.k); got != tt.want {
			t.Fatalf("subarraySum(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
		}
	}

}

/*
// @lcpr case=start
// [1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

*/
