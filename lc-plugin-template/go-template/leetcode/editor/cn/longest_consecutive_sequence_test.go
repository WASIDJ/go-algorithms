/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=30403
 *
 * [128] 最长连续序列
 */

package leetcode_solutions

import "testing"

// @lc code=start
func longestConsecutive(nums []int) int {
	// 思路：使用哈希表存储数组中的元素，然后遍历数组，对于每个元素，检查它的前一个元素是否存在，如果不存在，则说明当前元素是一个连续序列的起点，然后继续检查后续的元素是否存在，直到找到连续序列的终点，记录连续序列的长度，并更新最大长度。
	// 失败1:time limit exceeded 使用 map[int]struct{} 替代 map[int]bool 来节省空间和提高性能
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	maxLength := 0
	for num := range numSet {
		if _, exists := numSet[num-1]; !exists { // 如果前一个元素不存在，说明当前元素是一个连续序列的起点
			currentNum := num
			currentLength := 1

			for { // 继续检查后续的元素是否存在
				if _, exists := numSet[currentNum+1]; !exists {
					break
				}
				currentNum++
				currentLength++
			}

			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

// @lc code=end

func TestLongestConsecutiveSequence(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [100,4,200,1,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,3,7,2,5,8,4,6,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,2]\n
// @lcpr case=end

*/
