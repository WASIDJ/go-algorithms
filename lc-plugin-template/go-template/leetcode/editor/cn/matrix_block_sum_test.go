/*
 * @lc app=leetcode.cn id=1314 lang=golang
 * @lcpr version=30403
 *
 * [1314] 矩阵区域和
 */

package leetcode_solutions

import "testing"

// @lc code=start
func matrixBlockSum(mat [][]int, k int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return [][]int{}
	}
	m, n := len(mat), len(mat[0])
	// prefix sum with extra row/col
	ps := make([][]int, m+1)
	for i := range ps {
		ps[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			ps[i][j] = ps[i-1][j] + ps[i][j-1] - ps[i-1][j-1] + mat[i-1][j-1]
		}
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			r1 := max1(0, i-k)
			c1 := max1(0, j-k)
			r2 := min1(m-1, i+k)
			c2 := min1(n-1, j+k)
			// +1 for prefix indices
			res[i][j] = ps[r2+1][c2+1] - ps[r1][c2+1] - ps[r2+1][c1] + ps[r1][c1]
		}
	}
	return res
}

func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func TestMatrixBlockSum(t *testing.T) {
	tests := []struct {
		mat  [][]int
		k    int
		want [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1, [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2, [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}}},
	}
	for _, tc := range tests {
		got := matrixBlockSum(tc.mat, tc.k)
		if !equal(got, tc.want) {
			t.Fatalf("matrixBlockSum(%v, %d) = %v, want %v", tc.mat, tc.k, got, tc.want)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n1\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n2\n
// @lcpr case=end

*/
