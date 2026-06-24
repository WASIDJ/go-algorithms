/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=30403
 *
 * [238] 除了自身以外数组的乘积
 */

package leetcode_solutions

import (
	"reflect"
	"testing"
)

// @lc code=start
func productExceptSelf(nums []int) []int {
	n := len(nums)
	pm := make([]int, n)
	sm := make([]int, n)
	pm[0] = nums[0]
	for i := 1; i < n; i++ {
		pm[i] = nums[i] * pm[i-1]
	}
	sm[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sm[i] = nums[i] * sm[i+1]
	}
	res := make([]int, n)
	res[0] = sm[1]
	res[n-1] = pm[n-2]
	for i := 1; i < n-1; i++ {
		res[i] = pm[i-1] * sm[i+1]
	}

	return res
}

// @lc code=end

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"example1", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"example2", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("nums=%v: got=%v, want=%v", tt.nums, got, tt.want)
			}
		})
	}

}

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

*/
