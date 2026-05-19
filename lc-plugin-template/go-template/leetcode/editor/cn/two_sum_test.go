/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=30403
 *
 * [1] 两数之和
 */

package leetcode_solutions

import "testing"

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 创建一个map来存储数值和对应的索引
	numMap := make(map[int]int)
	// 遍历数组
	for i, num := range nums {
		// 计算需要找到的另一个数
		complement := target - num
		// 如果这个数已经在map中，说明找到了两个数
		if j, ok := numMap[complement]; ok {
			return []int{j, i}
		}
		// 将当前数和索引存入map
		numMap[num] = i
	}
	return nil
}

// @lc code=end

func TestTwoSum(t *testing.T) {
	// your test code here
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := twoSum(nums, target)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, result[i])
		}
	}
}

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/
