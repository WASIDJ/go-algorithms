/*
 * @lc app=leetcode.cn id=1094 lang=golang
 * @lcpr version=30403
 *
 * [1094] 拼车
 */

package leetcode_solutions

import "testing"

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	// 0 <= fromi < toi <= 1000 直接全建上,想要提升使用线段树
	diff := make([]int, 1000)
	for _, ele := range trips {
		diff = increment1(ele[0], ele[1], ele[2], diff)
	}
	res := result1(diff)
	for _, v := range res {
		if v > capacity {
			return false
		}
	}
	return true
}
func increment1(numPassengers, from, to int, diff []int) []int {
	diff[from] += numPassengers
	if to+1 < len(diff) {
		// 不应该加1 to 那一站是下车
		diff[to] -= numPassengers
	}
	return diff
}
func result1(diff []int) []int {
	res := make([]int, len(diff))
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = diff[i] + res[i-1]
	}
	return res
}

// @lc code=end

func TestCarPooling(t *testing.T) {
	tests := []struct {
		trips    [][]int
		capacity int
		expected bool
	}{
		{
			trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
			capacity: 4,
			expected: false,
		},
		{
			trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
			capacity: 5,
			expected: true,
		},
		{
			trips:    [][]int{{2, 1, 5}, {3, 5, 7}},
			capacity: 3,
			expected: true,
		},
	}

	for _, tt := range tests {
		result := carPooling(tt.trips, tt.capacity)
		if result != tt.expected {
			t.Errorf("carPooling(%v, %d) = %v, expected %v", tt.trips, tt.capacity, result, tt.expected)
		}
	}
}

/*
// @lcpr case=start
// [[2,1,5],[3,3,7]]\n4\n
// @lcpr case=end

// @lcpr case=start
// [[2,1,5],[3,3,7]]\n5\n
// @lcpr case=end

*/
