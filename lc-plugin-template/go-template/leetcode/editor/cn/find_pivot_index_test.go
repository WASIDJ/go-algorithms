/*
 * @lc app=leetcode.cn id=724 lang=golang
 * @lcpr version=30403
 *
 * [724] 寻找数组的中心下标
 */

package leetcode_solutions

import "testing"

// @lc code=start
func pivotIndex(nums []int) int {
	// 思路 ps
	n := len(nums)
	ps := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ps[i] = ps[i-1] + nums[i-1]
	}
	sum := ps[n]
	for i := 0; i < n; i++ {
		leftSum := ps[i]
		rightSum := sum - ps[i+1]
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

// @lc code=end

func TestFindPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{1, 7, 3, 6, 5, 6}, 3},
		{"example2", []int{1, 2, 3}, -1},
		{"example3", []int{2, 1, -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotIndex(tt.nums)
			if got != tt.want {
				t.Fatalf("pivotIndex(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}

/*
// @lcpr case=start
// [1,7,3,6,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,-1]\n
// @lcpr case=end

*/
