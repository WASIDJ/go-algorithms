/*
 * @lc app=leetcode.cn id=49 lang=golang
 * @lcpr version=30403
 *
 * [49] 字母异位词分组
 */

package leetcode_solutions

import (
	"fmt"
	"strconv"
	"testing"
)

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 思路：hashmap 第一位单词 first letter last letter
	hashMap := make(map[string][]string)
	for _, s := range strs {
		code := encode(s)
		hashMap[code] = append(hashMap[code], s)
	}
	// 将hashMap中的值转换为结果
	var result [][]string
	for _, group := range hashMap {
		result = append(result, group)
	}
	return result
}
func encode(s string) string {
	// 统计每个字母出现的次数
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	// 将统计结果转换为字符串
	var code string
	for i, c := range count {
		if c > 0 {
			code += string(rune('a'+i)) + strconv.Itoa(c)
		}
	}
	return code
}

// @lc code=end

func TestGroupAnagrams(t *testing.T) {
	// your test code here
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	result := groupAnagrams(strs)

	if len(result) != len(expected) {
		t.Errorf("Expected %d groups, got %d", len(expected), len(result))
	}

	for _, group := range result {
		found := false
		for _, exp := range expected {
			if len(group) == len(exp) {
				match := true
				for _, word := range group {
					exists := false
					for _, expWord := range exp {
						if word == expWord {
							exists = true
							break
						}
					}
					if !exists {
						match = false
						break
					}
				}
				if match {
					found = true
					break
				}
			}
		}
		if !found {
			t.Errorf("Unexpected group: %v", group)
		}
	}

	fmt.Println("Test passed!")
}

/*
// @lcpr case=start
// ["eat","tea","tan","ate","nat","bat"]\n
// @lcpr case=end

// @lcpr case=start
// [""]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/
