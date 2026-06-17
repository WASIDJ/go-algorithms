/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30403
 *
 * [3] 无重复字符的最长子串
 */

package leetcode_solutions

import "testing"

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 思路:滑动窗口
	// 什么时候扩大 需要更新什么
	windows := make(map[rune]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := rune(s[right])
		right++
		windows[c] = windows[c] + 1
		// 什么时候缩小 int 不大于2
		for windows[c] > 1 {
			d := rune(s[left])
			left++
			windows[d] = windows[d] - 1
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}

// @lc code=end

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
