/*
 * @lc app=leetcode.cn id=304 lang=golang
 * @lcpr version=30403
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

package leetcode_solutions

import "testing"

// @lc code=start
type NumMatrix struct {
	preNum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	if n == 0 {
		return NumMatrix{}
	}
	preNum := make([][]int, m+1)
	for i := range preNum {
		preNum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preNum[i][j] = preNum[i-1][j] + preNum[i][j-1] + matrix[i-1][j-1] - preNum[i-1][j-1]
		}
	}
	return NumMatrix{preNum: preNum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preNum[row2+1][col2+1] - this.preNum[row2+1][col1] - this.preNum[row1][col2+1] + this.preNum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

func TestRangeSumQuery2dImmutable(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	obj := Constructor(matrix)

	if got := obj.SumRegion(2, 1, 4, 3); got != 8 {
		t.Fatalf("SumRegion(2,1,4,3) = %d, want %d", got, 8)
	}
	if got := obj.SumRegion(1, 1, 2, 2); got != 11 {
		t.Fatalf("SumRegion(1,1,2,2) = %d, want %d", got, 11)
	}
	if got := obj.SumRegion(1, 2, 2, 4); got != 12 {
		t.Fatalf("SumRegion(1,2,2,4) = %d, want %d", got, 12)
	}

}

/*
// @lcpr case=start
// ["NumMatrix","sumRegion","sumRegion","sumRegion"]\n[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]\n
// @lcpr case=end

*/
