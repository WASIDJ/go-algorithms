/*
 * @lc app=leetcode.cn id=76 lang=golang
 * @lcpr version=30403
 *
 * [76] 最小覆盖子串
 */

package leetcode_solutions

import "testing"

// @lc code=start
func minWindow(s string, t string) string {
	// 思路:滑动窗口
	windows := make(map[byte]int)
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	valid := 0
	start, length := 0, 1<<31-1
	// 什么时候扩大 需要更新什么
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}
		// 什么时候缩小 需要更新什么
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if windows[d] == need[d] {
					valid--
				}
				windows[d]--
			}
		}
	}

	// 我需要什么
	if length == 1<<31-1 {
		return ""
	}
	return s[start : start+length]
}

// @lc code=end

func TestMinimumWindowSubstring(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// "ADOBECODEBANC"\n"ABC"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

*/
