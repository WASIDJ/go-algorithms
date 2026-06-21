/*
 * @lc app=leetcode.cn id=438 lang=golang
 * @lcpr version=30403
 *
 * [438] 找到字符串中所有字母异位词
 */

package leetcode_solutions

import "testing"

// @lc code=start
func findAnagrams(s string, p string) []int {
	// 感觉是和 3 很像 滑动窗口
	need := make(map[byte]int)
	windows := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	vaild := 0
	left, right := 0, 0
	res := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				vaild++
			}
		}
		for vaild == len(need) {
			if right-left == len(p) {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if windows[d] == need[d] {
					vaild--
				}
				windows[d]--
			}
		}
	}
	return res
}

// @lc code=end

func TestFindAllAnagramsInAString(t *testing.T) {
	s := "cbaebabacd"
	p := "abc"
	expected := []int{0, 6}
	result := findAnagrams(s, p)
	
	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
		return
	}
	
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %v, got %v", expected, result)
			return
		}
	}
}

/*
// @lcpr case=start
// "cbaebabacd"\n"abc"\n
// @lcpr case=end

// @lcpr case=start
// "abab"\n"ab"\n
// @lcpr case=end

*/
