/*
 * @lc app=leetcode.cn id=48 lang=golang
 * @lcpr version=30403
 *
 * [48] 旋转图像
 */

package leetcode_solutions

import "testing"

// @lc code=start
func rotate(matrix [][]int) {
	// 思路 对角线 再逐行reverse
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for _, ele := range matrix {
		reverse1(ele)
	}
}
func reverse1(ele []int) {
	l, r := 0, len(ele)-1
	for l < r {
		ele[l], ele[r] = ele[r], ele[l]
		l++
		r--
	}

}

// @lc code=end

func TestRotateImage(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]\n
// @lcpr case=end

*/
