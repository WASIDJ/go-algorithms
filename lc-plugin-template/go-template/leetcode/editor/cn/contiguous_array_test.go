/*
 * @lc app=leetcode.cn id=525 lang=golang
 * @lcpr version=30403
 *
 * [525] 连续数组
 */

package leetcode_solutions

import "testing"

// @lc code=start
func findMaxLength(nums []int) int {
	val2Idx := map[int]int{
		0: -1,
	}

	sum := 0
	res := 0

	for i, x := range nums {
		if x == 0 {
			sum--
		} else {
			sum++
		}

		if idx, ok := val2Idx[sum]; ok {
			res = max2(res, i-idx)
		} else {
			val2Idx[sum] = i
		}
	}

	return res
}
func max2(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

func TestContiguousArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "case1", nums: []int{0, 1}, want: 2},
		{name: "case2", nums: []int{0, 1, 0}, want: 2},
		{name: "case3", nums: []int{0, 1, 1, 1, 1, 1, 0, 0, 0}, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.nums); got != tt.want {
				t.Fatalf("findMaxLength(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}

}

/*
// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1,1,1,1,0,0,0]\n
// @lcpr case=end

*/
