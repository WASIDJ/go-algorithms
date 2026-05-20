/*
 * @lc app=leetcode.cn id=283 lang=golang
 * @lcpr version=30403
 *
 * [283] 移动零
 *
 * https://leetcode.cn/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (63.78%)
 * Likes:    2921
 * Dislikes: 0
 * Total Accepted:    2.4M
 * Total Submissions: 3.7M
 * Testcase Example:  '[0,1,0,3,12]\n[0]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 *
 * 示例 2:
 *
 * 输入: nums = [0]
 * 输出: [0]
 *
 *
 *
 * 提示:
 *
 *
 *
 * 1 <= nums.length <= 10^4
 * -2^31 <= nums[i] <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：你能尽量减少完成的操作次数吗？
 *
 */

package leetcode_solutions

import "testing"

// @lc code=start
func moveZeroes(nums []int) {
	//思路1：两个数组 一个 非零 一个零，最后合并
	//失败1:空间复杂度 O(n) 需要优化为 O(1)
	// 思路2: 双指针优化，使用一个指针遍历数组，另一个指针记录非零元素的位置，
	// 当遇到非零元素时，将其放到非零指针的位置，并将非零指针向前移动，
	// 最后将非零指针之后的元素全部置为0。
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}

}

// @lc code=end

func TestMoveZeroes(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [0,1,0,3,12]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
