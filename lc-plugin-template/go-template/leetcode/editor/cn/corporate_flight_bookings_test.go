/*
 * @lc app=leetcode.cn id=1109 lang=golang
 * @lcpr version=30403
 *
 * [1109] 航班预订统计
 */

package leetcode_solutions

import (
	"reflect"
	"testing"
)

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	// 思路差分数组
	diff := make([]int, n+1)
	for _, ele := range bookings {
		diff = increment(ele[0], ele[1], ele[2], diff)
	}
	res := result(diff, n+1)
	return res[1:]
}
func increment(i, j, val int, diff []int) []int {
	diff[i] += val
	if j+1 < len(diff) {
		diff[j+1] -= val
	}
	return diff
}
func result(diff []int, n int) []int {
	res := make([]int, n)
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = diff[i] + res[i-1]
	}
	return res
}

// @lc code=end

func TestCorporateFlightBookings(t *testing.T) {
	tests := []struct {
		bookings [][]int
		n        int
		want     []int
	}{
		{[][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5, []int{10, 55, 45, 25, 25}},
		{[][]int{{1, 2, 10}, {2, 2, 15}}, 2, []int{10, 25}},
	}
	for _, tt := range tests {
		got := corpFlightBookings(tt.bookings, tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("corpFlightBookings(%v, %d) = %v, want %v", tt.bookings, tt.n, got, tt.want)
		}
	}

}

/*
// @lcpr case=start
// [[1,2,10],[2,3,20],[2,5,25]]\n5\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,10],[2,2,15]]\n2\n
// @lcpr case=end

*/
