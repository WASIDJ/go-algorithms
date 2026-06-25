/*
 * @lc app=leetcode.cn id=59 lang=golang
 * @lcpr version=30403
 *
 * [59] 螺旋矩阵 II
 */

package leetcode_solutions

import (
	"reflect"
	"testing"
)

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	upper_bound := 0
	lower_bound := n - 1
	left_bound := 0
	right_bound := n - 1
	i := 1
	for i <= n*n {
		// 向右
		if upper_bound <= lower_bound {
			for j := left_bound; j <= right_bound; j++ {
				matrix[upper_bound][j] = i
				i++
			}
			upper_bound++
		}
		// 向下
		if left_bound <= right_bound {
			for j := upper_bound; j <= lower_bound; j++ {
				matrix[j][right_bound] = i
				i++
			}
			right_bound--
		}
		// 向左
		if upper_bound <= lower_bound {
			for j := right_bound; j >= left_bound; j-- {
				matrix[lower_bound][j] = i
				i++
			}
			lower_bound--
		}
		// 向上
		if left_bound <= right_bound {
			for j := lower_bound; j >= upper_bound; j-- {
				matrix[j][left_bound] = i
				i++
			}
			left_bound++
		}
	}
	return matrix
}

// @lc code=end

func TestSpiralMatrixIi(t *testing.T) {
	tests := []struct {
		n    int
		want [][]int
	}{
		{
			n:    1,
			want: [][]int{{1}},
		},
		{
			n:    3,
			want: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}},
		},
	}

	for _, tt := range tests {
		got := generateMatrix(tt.n)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("generateMatrix(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}

}

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
