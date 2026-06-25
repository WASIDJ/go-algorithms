/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=30403
 *
 * [54] 螺旋矩阵
 */

package leetcode_solutions

import (
	"reflect"
	"testing"
)

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	var res []int
	upper_bound := 0
	lower_bound := m - 1
	left_bound := 0
	right_bound := n - 1

	for len(res) < m*n {
		// 向右
		if upper_bound <= lower_bound {
			for i := left_bound; i <= right_bound; i++ {
				res = append(res, matrix[upper_bound][i])
			}
			upper_bound++
		}
		// 向下
		if left_bound <= right_bound {
			for i := upper_bound; i <= lower_bound; i++ {
				res = append(res, matrix[i][right_bound])
			}
			right_bound--
		}
		// 向左
		if upper_bound <= lower_bound {
			for i := right_bound; i >= left_bound; i-- {
				res = append(res, matrix[lower_bound][i])
			}
			lower_bound--
		}
		// 向上
		if left_bound <= right_bound {
			for i := lower_bound; i >= upper_bound; i-- {
				res = append(res, matrix[i][left_bound])
			}
			left_bound++
		}
	}
	return res

}

// @lc code=end

func TestSpiralMatrix(t *testing.T) {
	cases := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{"3x4", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := spiralOrder(tc.matrix)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("spiralOrder(%v) = %v, want %v", tc.matrix, got, tc.want)
			}
		})
	}

}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/
